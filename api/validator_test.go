package api

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestValidCurrency(t *testing.T) {
	testCases := []struct {
		name     string
		currency string
		valid    bool
	}{
		{
			name:     "Valid USD",
			currency: "USD",
			valid:    true,
		},
		{
			name:     "Valid EUR",
			currency: "EUR",
			valid:    true,
		},
		{
			name:     "Valid INR",
			currency: "INR",
			valid:    true,
		},
		{
			name:     "Invalid Currency",
			currency: "XYZ",
			valid:    false,
		},
		{
			name:     "Empty Currency",
			currency: "",
			valid:    false,
		},
		{
			name:     "Lowercase Currency",
			currency: "usd",
			valid:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a validator instance
			v := validator.New()
			v.RegisterValidation("currency", validCurrency)

			// Test struct with currency field
			type TestStruct struct {
				Currency string `validate:"currency"`
			}

			testStruct := TestStruct{Currency: tc.currency}
			err := v.Struct(testStruct)

			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestValidCurrencyWithNonStringField(t *testing.T) {
	// Test that validator returns false for non-string fields
	v := validator.New()
	v.RegisterValidation("currency", validCurrency)

	type TestStruct struct {
		Amount int64 `validate:"currency"`
	}

	testStruct := TestStruct{Amount: 100}
	err := v.Struct(testStruct)

	// Should return validation error because the field is not a string
	require.Error(t, err)
}

