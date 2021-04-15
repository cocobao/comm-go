package httpex

import (
	"compress/gzip"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cocobao/comm-go/log"
	"golang.org/x/net/context/ctxhttp"
)

var (
	cookies []*http.Cookie
)

func GetCookies() []*http.Cookie {
	return cookies
}

func AddCookies(c *http.Cookie) {
	cookies = append(cookies, c)
}

func GetRequestWithCookies(url string, headEx map[string]string, retHandle func(res *http.Response, bd []byte) error) error {
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

	if cookies != nil && len(cookies) > 0 {
		for _, v := range cookies {
			req.AddCookie(v)
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

	if res.Header.Get("Content-Encoding") == "gzip" {
		render, err := gzip.NewReader(res.Body)
		if err != nil {
			log.Error(err)
			return err
		} else {
			defer render.Close()

			result, err = ioutil.ReadAll(render)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	} else {
		result, err = ioutil.ReadAll(res.Body)
		if err != nil {
			log.Info(string(result))
		}
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("StatusCode:%d", res.StatusCode)
	}

	if retHandle != nil {
		return retHandle(res, result)
	}

	return nil
}
