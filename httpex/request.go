package httpex

import (
	"net/http"
)

func setHeaders(req *http.Request) {
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
}

func GetRequestHandle(url string, headEx map[string]string, preHandle func(bd []byte) ([]byte, error)) error {
	return GetRequest(url, headEx, preHandle, nil)
}
