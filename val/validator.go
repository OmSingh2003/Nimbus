package val

import (
	"fmt"
	"regexp"
)

var (
	isValidateUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidateEmail    = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`).MatchString
)

func validateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := validateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidateUsername(value) {
		return fmt.Errorf("must contain only letters, digits, or underscore")
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := validateString(value, 3, 200); err != nil {
		return err
	}
	if !isValidateEmail(value) {
		return fmt.Errorf("must be a valid email address")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := validateString(value, 3, 100); err != nil {
		return err
	}
	// Full name should only contain letters and spaces
	for _, char := range value {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == ' ') {
			return fmt.Errorf("must contain only letters and spaces")
		}
	}
	return nil
}

func ValidatePassword(value string) error {
	if err := validateString(value, 6, 100); err != nil {
		return err
	}
	// For passwords 8 chars or longer, require lowercase and digit
	if len(value) >= 8 {
		hasLower := false
		hasDigit := false
		for _, char := range value {
			if char >= 'a' && char <= 'z' {
				hasLower = true
			}
			if char >= '0' && char <= '9' {
				hasDigit = true
			}
		}
		if !hasLower {
			return fmt.Errorf("password must contain at least one lowercase letter")
		}
		if !hasDigit {
			return fmt.Errorf("password must contain at least one digit")
		}
	}
	return nil
}

func ValidateSecretKey(value string) error {
	return validateString(value, 32, 128)
}

func ValidateCurrency(value string) error {
	if err := validateString(value, 3, 3); err != nil {
		return err
	}
	// Currency should be uppercase letters only
	for _, char := range value {
		if !(char >= 'A' && char <= 'Z') {
			return fmt.Errorf("must contain only uppercase letters")
		}
	}
	return nil
}

func ValidateAmount(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be positive")
	}
	if value > 1000000000 {
		return fmt.Errorf("must not exceed 1,000,000,000")
	}
	return nil
}

func ValidateID(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be positive")
	}
	return nil
}

func ValidateAccountOwner(value string) error {
	return ValidateUsername(value)
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	if err := validateString(value, 32, 128); err != nil {
		return err
	}
	return nil
}

func ValidatePageNumber(value int32) error {
	if value <= 0 {
		return fmt.Errorf("must be positive")
	}
	return nil
}

func ValidatePageSize(value int32) error {
	if value <= 0 {
		return fmt.Errorf("must be positive")
	}
	if value > 100 {
		return fmt.Errorf("must not exceed 100")
	}
	return nil
}
