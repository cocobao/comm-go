package utility

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrCipherKey           = errors.New("The secret key is wrong and cannot be decrypted. Please check")
	ErrKeyLengthSixteen    = errors.New("a sixteen or twenty-four or thirty-two length secret key is required")
	ErrKeyLengtheEight     = errors.New("a eight-length secret key is required")
	ErrKeyLengthTwentyFour = errors.New("a twenty-four-length secret key is required")
	ErrPaddingSize         = errors.New("padding size error please check the secret key or iv")
	ErrIvAes               = errors.New("a sixteen-length ivaes is required")
	ErrIvDes               = errors.New("a eight-length ivdes key is required")
)

func PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func PKCS5UnPadding(plainText []byte) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number > length {
		return nil, ErrPaddingSize
	}
	return plainText[:length-number], nil
}

// encrypt
func AesCbcEncrypt(plainText, key []byte, ivAes []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	paddingText := PKCS5Padding(plainText, block.BlockSize())

	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != 16 {
			return nil, ErrIvAes
		} else {
			iv = ivAes
		}
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

// decrypt
func AesCbcDecrypt(cipherText, key []byte, ivAes []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				fmt.Println("runtime err:", err, "Check that the key or text is correct")
			default:
				fmt.Println("error:", err)
			}
		}
	}()
	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != 16 {
			return nil, ErrIvAes
		} else {
			iv = ivAes
		}
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)

	plainText, err := PKCS5UnPadding(paddingText)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
