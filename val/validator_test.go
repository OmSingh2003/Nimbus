package val

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateString(t *testing.T) {
	testCases := []struct {
		name      string
		value     string
		minLength int
		maxLength int
		expectErr bool
	}{
		{
			name:      "valid string",
			value:     "test",
			minLength: 3,
			maxLength: 10,
			expectErr: false,
		},
		{
			name:      "too short",
			value:     "ab",
			minLength: 3,
			maxLength: 10,
			expectErr: true,
		},
		{
			name:      "too long",
			value:     "this is way too long",
			minLength: 3,
			maxLength: 10,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateString(tc.value, tc.minLength, tc.maxLength)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	testCases := []struct {
		name      string
		username  string
		expectErr bool
	}{
		{
			name:      "valid username",
			username:  "test_user123",
			expectErr: false,
		},
		{
			name:      "valid short username",
			username:  "abc",
			expectErr: false,
		},
		{
			name:      "too short",
			username:  "ab",
			expectErr: true,
		},
		{
			name:      "contains uppercase",
			username:  "TestUser",
			expectErr: true,
		},
		{
			name:      "contains special chars",
			username:  "test@user",
			expectErr: true,
		},
		{
			name:      "contains space",
			username:  "test user",
			expectErr: true,
		},
		{
			name:      "empty string",
			username:  "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateUsername(tc.username)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateFullName(t *testing.T) {
	testCases := []struct {
		name      string
		fullName  string
		expectErr bool
	}{
		{
			name:      "valid full name",
			fullName:  "John Doe",
			expectErr: false,
		},
		{
			name:      "valid single name",
			fullName:  "John",
			expectErr: false,
		},
		{
			name:      "valid multiple names",
			fullName:  "John Michael Doe",
			expectErr: false,
		},
		{
			name:      "too short",
			fullName:  "Jo",
			expectErr: true,
		},
		{
			name:      "contains numbers",
			fullName:  "John123",
			expectErr: true,
		},
		{
			name:      "contains special chars",
			fullName:  "John@Doe",
			expectErr: true,
		},
		{
			name:      "empty string",
			fullName:  "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateFullName(tc.fullName)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		name      string
		email     string
		expectErr bool
	}{
		{
			name:      "valid email",
			email:     "test@example.com",
			expectErr: false,
		},
		{
			name:      "valid email with subdomain",
			email:     "user@mail.example.com",
			expectErr: false,
		},
		{
			name:      "valid email with plus",
			email:     "user+tag@example.com",
			expectErr: false,
		},
		{
			name:      "invalid missing @",
			email:     "testexample.com",
			expectErr: true,
		},
		{
			name:      "invalid missing domain",
			email:     "test@",
			expectErr: true,
		},
		{
			name:      "invalid missing username",
			email:     "@example.com",
			expectErr: true,
		},
		{
			name:      "too short",
			email:     "a@",
			expectErr: true,
		},
		{
			name:      "empty string",
			email:     "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateEmail(tc.email)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		name      string
		password  string
		expectErr bool
	}{
		{
			name:      "valid password",
			password:  "password123",
			expectErr: false,
		},
		{
			name:      "valid minimum length",
			password:  "123456",
			expectErr: false,
		},
		{
			name:      "too short",
			password:  "12345",
			expectErr: true,
		},
		{
			name:      "empty string",
			password:  "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePassword(tc.password)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateSecretKey(t *testing.T) {
	testCases := []struct {
		name      string
		secretKey string
		expectErr bool
	}{
		{
			name:      "valid secret key",
			secretKey: "abcdefghijklmnopqrstuvwxyz123456",
			expectErr: false,
		},
		{
			name:      "too short",
			secretKey: "short",
			expectErr: true,
		},
		{
			name:      "empty string",
			secretKey: "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateSecretKey(tc.secretKey)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateCurrency(t *testing.T) {
	testCases := []struct {
		name      string
		currency  string
		expectErr bool
	}{
		{
			name:      "valid USD",
			currency:  "USD",
			expectErr: false,
		},
		{
			name:      "valid EUR",
			currency:  "EUR",
			expectErr: false,
		},
		{
			name:      "valid JPY",
			currency:  "JPY",
			expectErr: false,
		},
		{
			name:      "lowercase",
			currency:  "usd",
			expectErr: true,
		},
		{
			name:      "too short",
			currency:  "US",
			expectErr: true,
		},
		{
			name:      "too long",
			currency:  "USDD",
			expectErr: true,
		},
		{
			name:      "contains numbers",
			currency:  "US1",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateCurrency(tc.currency)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateAmount(t *testing.T) {
	testCases := []struct {
		name      string
		amount    int64
		expectErr bool
	}{
		{
			name:      "valid amount",
			amount:    100,
			expectErr: false,
		},
		{
			name:      "large valid amount",
			amount:    999999999,
			expectErr: false,
		},
		{
			name:      "zero amount",
			amount:    0,
			expectErr: true,
		},
		{
			name:      "negative amount",
			amount:    -100,
			expectErr: true,
		},
		{
			name:      "too large amount",
			amount:    1000000001,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateAmount(tc.amount)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateID(t *testing.T) {
	testCases := []struct {
		name      string
		id        int64
		expectErr bool
	}{
		{
			name:      "valid ID",
			id:        1,
			expectErr: false,
		},
		{
			name:      "large valid ID",
			id:        999999999,
			expectErr: false,
		},
		{
			name:      "zero ID",
			id:        0,
			expectErr: true,
		},
		{
			name:      "negative ID",
			id:        -1,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateID(tc.id)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidatePasswordEnhanced(t *testing.T) {
	testCases := []struct {
		name      string
		password  string
		expectErr bool
	}{
		{
			name:      "simple valid password",
			password:  "simple",
			expectErr: false,
		},
		{
			name:      "strong password",
			password:  "Password123!",
			expectErr: false,
		},
		{
			name:      "8+ chars without lowercase",
			password:  "PASSWORD123",
			expectErr: true,
		},
		{
			name:      "8+ chars without digits",
			password:  "Password",
			expectErr: true,
		},
		{
			name:      "valid 8+ chars with requirements",
			password:  "password123",
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePassword(tc.password)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateAccountOwner(t *testing.T) {
	testCases := []struct {
		name      string
		owner     string
		expectErr bool
	}{
		{
			name:      "valid owner",
			owner:     "test_user",
			expectErr: false,
		},
		{
			name:      "invalid owner format",
			owner:     "Test User",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateAccountOwner(tc.owner)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// Benchmark tests
func BenchmarkValidateUsername(b *testing.B) {
	username := "test_user123"
	for i := 0; i < b.N; i++ {
		_ = ValidateUsername(username)
	}
}

func BenchmarkValidateEmail(b *testing.B) {
	email := "test@example.com"
	for i := 0; i < b.N; i++ {
		_ = ValidateEmail(email)
	}
}

func BenchmarkValidateFullName(b *testing.B) {
	fullName := "John Doe"
	for i := 0; i < b.N; i++ {
		_ = ValidateFullName(fullName)
	}
}

func BenchmarkValidatePassword(b *testing.B) {
	password := "password123"
	for i := 0; i < b.N; i++ {
		_ = ValidatePassword(password)
	}
}

func BenchmarkValidateCurrency(b *testing.B) {
	currency := "USD"
	for i := 0; i < b.N; i++ {
		_ = ValidateCurrency(currency)
	}
}

