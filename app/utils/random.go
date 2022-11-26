package utils

import (
	"math/rand"
	"time"
)

var runes = []rune("01213456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-")

func RandomURL(size int) string {
	str := make([]rune, size)

	rand.Seed(time.Now().UnixNano())

	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}
