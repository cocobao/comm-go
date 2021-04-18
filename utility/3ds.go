package utility

import (
	"encoding/base64"

	"github.com/wumansgy/goEncrypt"
)

func Decrypt3DES(text, key, iv string) (string, error) {
	textData, _ := base64.StdEncoding.DecodeString(text)
	key = Md5(key)
	iv = Md5(iv)
	result, err := goEncrypt.TripleDesDecrypt(textData, []byte(key)[:24], []byte(iv)[:8]...)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func Encrypt3DES(text []byte, key, iv string) (string, error) {
	key = Md5(key)
	iv = Md5(iv)

	result, err := goEncrypt.TripleDesEncrypt(text, []byte(key)[:24], []byte(iv)[:8]...)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(result), nil
}
