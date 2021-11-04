package utility

import (
	"github.com/afocus/captcha"
	"github.com/mojocn/base64Captcha"
)

var (
	Cap *captcha.Captcha

	capstore = &CaptchaStore{}
)

//获取图片验证码
func GetCaptcha(width, height, length int) (id string, b64s string, err error) {
	driver := &base64Captcha.DriverDigit{
		Height:   width,
		Width:    height,
		Length:   length,
		MaxSkew:  0.1,
		DotCount: 15,
	}

	c := base64Captcha.NewCaptcha(driver, capstore)
	return c.Generate()
}
