package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	k := len(chars)

	for i := 0; i < length; i++ {
		c := chars[rand.Intn(k)]

		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwnerName() string {
	return RandomString(10)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, BRL}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
