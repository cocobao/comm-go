package utility

import (
	"github.com/skip2/go-qrcode"
)

func GetQrcode(str string) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(str, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
