package utility

import (
	"time"

	"github.com/cocobao/comm-go/service"
)

type CaptchaStore struct {
}

func (c *CaptchaStore) Get(id string, clear bool) string {
	key := "captcha_" + id

	s, _ := service.CacheGet(key)
	if clear {
		service.CacheDel(key)
	}
	return s
}

func (c *CaptchaStore) Set(id string, value string) {
	service.CacheSet("captcha_"+id, value, time.Minute)
}

func (c *CaptchaStore) Verify(id, answer string, clear bool) bool {
	return (c.Get(id, clear) == answer)
}
