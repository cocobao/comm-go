package utility

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	h := md5.New()
	b := []byte(str)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
