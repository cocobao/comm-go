package utility

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	letterBytes0 = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=")
	letterBytes1 = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	letterBytes2 = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letterBytes3 = []rune("1234567890")
)

func GetRandNumCode6() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

func RandString0(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterBytes0[rand.Intn(len(letterBytes0))]
	}
	return string(b)
}

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterBytes1[rand.Intn(len(letterBytes1))]
	}
	return string(b)
}

func RandPureLetterString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterBytes2[rand.Intn(len(letterBytes2))]
	}
	return string(b)
}

func RandInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func RandPureNumberString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterBytes3[rand.Intn(len(letterBytes3))]
	}
	return string(b)
}
