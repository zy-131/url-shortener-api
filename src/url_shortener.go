package main

import (
	"math/rand"
	"time"
)

var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateShortURL() string {
	length := 6
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}

	return string(b)
}
