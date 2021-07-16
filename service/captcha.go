package service

import (
	"time"
)

type CaptchaStore struct {
}

func (c *CaptchaStore) Get(id string, clear bool) string {
	key := "cocobao_service_captcha_" + id

	s, _ := CacheGet(key)
	if clear {
		CacheDel(key)
	}
	return s
}

func (c *CaptchaStore) Set(id string, value string) {
	CacheSet("cocobao_service_captcha_"+id, value, time.Minute*10)
}

func (c *CaptchaStore) Verify(id, answer string, clear bool) bool {
	return (c.Get(id, clear) == answer)
}
