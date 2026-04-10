// Package Util
package util

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt Generate a random integer between given min and max
func RandomInt(min, max int64) int64 {
	return int64(min) + rand.Int64N(max-min+1)
}

// RandomString generate a random string of given length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for range n {
		c := alphabet[rand.Int64N(int64(k))]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generate a random name of the owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency genreate a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := int64(len(currencies))
	return currencies[rand.Int64N(n)]
}

// RandomEmail genreate a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
