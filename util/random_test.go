package util

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
    min := int64(1)
    max := int64(1000)
    n := RandomInt(min, max)
    require.GreaterOrEqual(t, n, min)
    require.LessOrEqual(t, n, max)
}

func TestRandomString(t *testing.T) {
    n := 6
    s := RandomString(n)
    require.Len(t, s, n)
}

func TestRandomOwner(t *testing.T) {
    owner := RandomOwner()
    require.Len(t, owner, 6)
}

func TestRandomMoney(t *testing.T) {
    money := RandomMoney()
    require.GreaterOrEqual(t, money, int64(0))
    require.LessOrEqual(t, money, int64(1000))
}

func TestRandomCurrency(t *testing.T) {
	currency := RandomCurrency()
	require.Contains(t, []string{"USD", "EUR", "INR"}, currency)
}

func TestRandomStrongPassword(t *testing.T) {
	password := RandomStrongPassword()
	require.NotEmpty(t, password)
	
	// Should pass password validation
	err := ValidatePassword(password)
	require.NoError(t, err)
	
	// Should be hashable
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
}
