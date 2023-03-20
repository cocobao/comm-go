package httpex

import (
	"compress/flate"
	"compress/gzip"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"comm-go/log"
	"comm-go/utility"

	"golang.org/x/net/context/ctxhttp"
)

func GetNetFile(href, savepath string) (string, error) {
	time.Sleep(time.Second)
	req, err := http.NewRequest("GET", href, nil)
	if err != nil {
		log.Debug("NewRequest err:", err.Error())
		return "", err
	}
	setHeaders(req)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	ctxto, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	res, err := ctxhttp.Do(ctxto, client, req)
	if err != nil {
		log.Debug("req err:", err)
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		log.Debugf("get result:%d\n", res.StatusCode)
		return "", fmt.Errorf("result err:%d", res.StatusCode)
	}

	filePath := savepath
	if s := res.Header.Get("Content-Disposition"); s != "" {
		ss := strings.Split(s, "filename=")[1]
		ss, _ = url.QueryUnescape(ss)
		log.Debugf("download file name:%v", ss)

		filePath = path.Join(filePath, ss)
	}

	if utility.IsExist(filePath) {
		return filePath, nil
	} else {
		os.MkdirAll(path.Dir(filePath), os.ModePerm)
	}

	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			log.Debug("reader gzip fail,", err)
			return "", err
		}
	case "deflate":
		reader = flate.NewReader(res.Body)
	default:
		reader = res.Body
	}
	defer res.Body.Close()

	f, err := os.Create(filePath)
	if err != nil {
		log.Debug("creat file fail, err:", err)
		return "", err
	}
	defer f.Close()
	if _, err = io.Copy(f, reader); err != nil {
		log.Debug("copy fail,", err)
		return "", err
	}
	return filePath, nil
}
