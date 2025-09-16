package testutil

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// RandomString generates a random string of n characters
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}

func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "MYR"}
	return currencies[rng.Intn(len(currencies))]
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6) // e.g. "xqjzpa"
}

// RandomMoney generates a random balance amount
func RandomMoney() int64 {
	return rng.Int63n(1000) + 1 // between 1 and 1000
}
