package utility

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/cocobao/comm-go/log"
	"golang.org/x/net/context/ctxhttp"
)

func setHeaders(req *http.Request) {
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json, text/javascript, */*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Set("Content-type", "text/plain")
}

func GetRequestHandle(url string, headEx map[string]string, preHandle func(bd []byte) ([]byte, error)) error {
	return GetRequest(url, headEx, preHandle, nil)
}

func GetRequest(url string, headEx map[string]string, preHandle func(bd []byte) ([]byte, error), retData interface{}) error {
	var req *http.Request
	var res *http.Response
	var err error

	tryTime := 0

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

tryAgain:
	time.Sleep(time.Second)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Debug("NewRequest err:", err.Error())
		return err
	}
	setHeaders(req)
	if headEx != nil {
		for k, v := range headEx {
			req.Header.Set(k, v)
		}
	}

	ctx := context.Background()
	ctxto, cancel := context.WithTimeout(ctx, 10*time.Minute)
	res, err = ctxhttp.Do(ctxto, client, req)
	defer cancel()
	if err != nil {
		log.Warn("http get err:", err, tryTime)
		select {
		case <-ctx.Done():
		default:
		}

		tryTime++
		if tryTime < 5 {
			goto tryAgain
		}
		return err
	}

	var result []byte
	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Info(string(result))
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("StatusCode:%d", res.StatusCode)
	}

	if preHandle != nil {
		result, err = preHandle(result)
		if err != nil {
			return err
		}
	}

	if retData != nil {
		err = json.Unmarshal(result, retData)
		if err != nil {
			tryTime++
			if tryTime < 5 {
				goto tryAgain
			}
		}
	}

	return nil
}

func PostRequst(url string, headEx map[string]string, bd []byte, retData interface{}) error {
	var req *http.Request
	var res *http.Response
	var err error
	tryTime := 0

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

tryAgain:
	time.Sleep(time.Second)
	req, err = http.NewRequest("POST", url, bytes.NewReader(bd))
	if err != nil {
		log.Debug("NewRequest err:", err.Error())
		return err
	}
	setHeaders(req)
	if headEx != nil {
		for k, v := range headEx {
			req.Header.Set(k, v)
		}
	}

	ctx := context.Background()
	ctxto, cancel := context.WithTimeout(ctx, 10*time.Minute)
	res, err = ctxhttp.Do(ctxto, client, req)
	defer cancel()
	if err != nil {
		log.Warn("push post err:", err, tryTime)
		select {
		case <-ctx.Done():
			return err
		default:
		}

		tryTime++
		if tryTime < 3 {
			goto tryAgain
		}
		return err
	}

	var result []byte
	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Info(string(result))
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("StatusCode:%d", res.StatusCode)
	}

	if retData != nil {
		return json.Unmarshal(result, retData)
	}
	return nil
}

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

	if IsExist(filePath) {
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

func PostFormData(href string, bd map[string]string, retData interface{}) error {
	payload := url.Values{}
	for k, v := range bd {
		payload.Set(k, v)
	}
	tryTime := 0
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
tryAgain:
	time.Sleep(time.Second)
	req, err := http.NewRequest("POST", href, strings.NewReader(payload.Encode()))
	if err != nil {
		log.Debug("NewRequest err:", err.Error())
		return err
	}
	setHeaders(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	ctx := context.Background()
	ctxto, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	res, err := ctxhttp.Do(ctxto, client, req)
	defer cancel()
	if err != nil {
		log.Warn("push post err:", err, tryTime)
		select {
		case <-ctx.Done():
			return err
		default:
		}

		tryTime++
		if tryTime < 3 {
			goto tryAgain
		}
		return err
	}

	var result []byte
	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Info(string(result))
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("StatusCode:%d", res.StatusCode)
	}

	if retData != nil {
		return json.Unmarshal(result, retData)
	}
	return nil
}

func PostForm(href string, bd map[string]string) ([]byte, error) {
	payload := url.Values{}
	for k, v := range bd {
		payload.Set(k, v)
	}
	tryTime := 0
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
tryAgain:
	req, err := http.NewRequest("POST", href, strings.NewReader(payload.Encode()))
	if err != nil {
		log.Debug("NewRequest err:", err.Error())
		return nil, err
	}
	setHeaders(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	ctx := context.Background()
	ctxto, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	res, err := ctxhttp.Do(ctxto, client, req)
	defer cancel()
	if err != nil {
		log.Warn("push post err:", err, tryTime)
		select {
		case <-ctx.Done():
			return nil, err
		default:
		}

		tryTime++
		if tryTime < 3 {
			time.Sleep(time.Second)
			goto tryAgain
		}
		return nil, err
	}

	var result []byte
	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Info(string(result))
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode:%d", res.StatusCode)
	}

	return result, nil
}
