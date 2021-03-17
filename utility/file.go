package utility

import (
	"io/ioutil"
	"os"
	"path"
)

func WriteFile(filePath string, content []byte) error {
	_ = os.MkdirAll(path.Dir(filePath), os.ModePerm)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	return err
}

func ReadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
