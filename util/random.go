package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano()) // Fixed: capitalized Now()
}

// Generates random integers between max and min
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] // Fixed: indexing alphabet with []
		sb.WriteByte(c)
	}
	return sb.String()
}

// generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{EUR, USD ,INR}
	n := len(currencies)
	return currencies[rand.Intn(n)] // Fixed: indexing currencies with []
}
// generates a random email 
func RandomEmail() string{
	return fmt.Sprintf("%s@gmail.com",RandomString(6))
}

// generates a random strong password for testing
func RandomStrongPassword() string {
	return fmt.Sprintf("%s%s%d!", 
		strings.ToUpper(RandomString(4)), 
		strings.ToLower(RandomString(4)), 
		RandomInt(10, 99))
}

// generates a random bank account number
func RandomAccountNumber() string {
	return fmt.Sprintf("%010d", RandomInt(1000000000, 9999999999))
}
