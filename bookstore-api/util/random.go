package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnoprstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomBookName() string {
	return RandomString(6)
}

func RandomBookPublication() string {
	return RandomString(6)
}
