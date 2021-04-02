package httpex

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cocobao/comm-go/log"
	"golang.org/x/net/context/ctxhttp"
)

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
