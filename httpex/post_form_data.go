package httpex

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cocobao/comm-go/log"
	"golang.org/x/net/context/ctxhttp"
)

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
