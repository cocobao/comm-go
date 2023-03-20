package httpex

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"comm-go/log"

	"golang.org/x/net/context/ctxhttp"
)

func DeleteRequest(url string, headEx map[string]string, retData interface{}) error {
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
	req, err = http.NewRequest(http.MethodDelete, url, nil)
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

	if retData != nil {
		err = json.Unmarshal(result, retData)
	}

	return nil
}
