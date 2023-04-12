package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var newRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int64) int64 {
	return min + newRand.Int63n(max-min+1)
}

func RandString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[newRand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandString(6)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "EGP"}
	return currencies[newRand.Intn(len(currencies))]
}
