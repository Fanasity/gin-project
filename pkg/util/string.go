package util

import (
	"math/rand"
	"time"
)

func GenerateCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
