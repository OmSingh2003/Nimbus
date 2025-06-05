package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSupportedCurrency(t *testing.T) {
	testCases := []struct {
		name     string
		currency string
		want     bool
	}{
		{
			name:     "USD is supported",
			currency: USD,
			want:     true,
		},
		{
			name:     "EUR is supported",
			currency: EUR,
			want:     true,
		},
		{
			name:     "INR is supported",
			currency: INR,
			want:     true,
		},
		{
			name:     "USD string literal is supported",
			currency: "USD",
			want:     true,
		},
		{
			name:     "EUR string literal is supported",
			currency: "EUR",
			want:     true,
		},
		{
			name:     "INR string literal is supported",
			currency: "INR",
			want:     true,
		},
		{
			name:     "Unsupported currency XYZ",
			currency: "XYZ",
			want:     false,
		},
		{
			name:     "Unsupported currency JPY",
			currency: "JPY",
			want:     false,
		},
		{
			name:     "Empty string is not supported",
			currency: "",
			want:     false,
		},
		{
			name:     "Lowercase usd is not supported",
			currency: "usd",
			want:     false,
		},
		{
			name:     "Mixed case Usd is not supported",
			currency: "Usd",
			want:     false,
		},
		{
			name:     "Currency with spaces is not supported",
			currency: " USD ",
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsSupportedCurrency(tc.currency)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestCurrencyConstants(t *testing.T) {
	// Test that currency constants have expected values
	require.Equal(t, "USD", USD)
	require.Equal(t, "EUR", EUR)
	require.Equal(t, "INR", INR)

	// Test that all constants are supported
	require.True(t, IsSupportedCurrency(USD))
	require.True(t, IsSupportedCurrency(EUR))
	require.True(t, IsSupportedCurrency(INR))
}

