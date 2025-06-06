package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	validPassword := "Passw0rd!"
	hashedPassword, err := HashPassword(validPassword)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(validPassword, hashedPassword)
	require.NoError(t, err)

	wrongPassword := "WrongPass1!"
	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

func TestValidatePassword(t *testing.T) {
	cases := []struct {
		password string
		expects  string
	}{
		{"short", "password must be at least 8 characters long"},
		{"VeryVeryVeryLongPassword1!ThatDefinitelyExceedsTheLimitOf64CharactersInLength", "password must be at most 64 characters long"},
		{"nocapitals1!", "password must include upper, lower, number, and special character"},
		{"NOLOWERCASE1!", "password must include upper, lower, number, and special character"},
		{"NoNumbers!!", "password must include upper, lower, number, and special character"},
		{"NoSpecials1", "password must include upper, lower, number, and special character"},
	}

	for _, c := range cases {
		err := ValidatePassword(c.password)
		require.EqualError(t, err, c.expects)
	}

	validPassword := "StrongPass1!"
	err := ValidatePassword(validPassword)
	require.NoError(t, err)
}

func TestHashPasswordWithInvalidInput(t *testing.T) {
	// Test hashing with invalid passwords
	invalidPasswords := []string{
		"short",
		"nocapitals1!",
		"NOLOWERCASE1!",
		"NoNumbers!!",
		"NoSpecials1",
	}
	
	for _, password := range invalidPasswords {
		_, err := HashPassword(password)
		require.Error(t, err)
	}
}

func TestCheckPasswordEdgeCases(t *testing.T) {
	validPassword := "TestPass1!"
	hashedPassword, err := HashPassword(validPassword)
	require.NoError(t, err)
	
	// Test with empty password
	err = CheckPassword("", hashedPassword)
	require.Error(t, err)
	
	// Test with empty hash
	err = CheckPassword(validPassword, "")
	require.Error(t, err)
	
	// Test with malformed hash
	err = CheckPassword(validPassword, "invalid-hash")
	require.Error(t, err)
}

func TestPasswordUniqueness(t *testing.T) {
	password := "SamePass1!"
	
	// Generate multiple hashes of the same password
	hash1, err := HashPassword(password)
	require.NoError(t, err)
	
	hash2, err := HashPassword(password)
	require.NoError(t, err)
	
	// Hashes should be different (due to salt)
	require.NotEqual(t, hash1, hash2)
	
	// But both should verify correctly
	err = CheckPassword(password, hash1)
	require.NoError(t, err)
	
	err = CheckPassword(password, hash2)
	require.NoError(t, err)
}

func TestUpdatePassword(t *testing.T) {
	oldPassword := "OldPass1!"
	newPassword := "NewPass1!"
	
	// Hash the old password
	oldHash, err := HashPassword(oldPassword)
	require.NoError(t, err)
	
	// Update password successfully
	newHash, err := UpdatePassword(oldPassword, newPassword, oldHash)
	require.NoError(t, err)
	require.NotEmpty(t, newHash)
	require.NotEqual(t, oldHash, newHash)
	
	// Verify new password works
	err = CheckPassword(newPassword, newHash)
	require.NoError(t, err)
	
	// Verify old password no longer works with new hash
	err = CheckPassword(oldPassword, newHash)
	require.Error(t, err)
	
	// Test with wrong old password
	_, err = UpdatePassword("WrongOld1!", newPassword, oldHash)
	require.Error(t, err)
	require.Contains(t, err.Error(), "current password is incorrect")
	
	// Test with invalid new password
	_, err = UpdatePassword(oldPassword, "weak", oldHash)
	require.Error(t, err)
	require.Contains(t, err.Error(), "password must be at least")
}
