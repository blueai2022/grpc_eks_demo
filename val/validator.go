package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateLength(value string, min, max int) error {
	n := len(value)

	if n < min || n > max {
		return fmt.Errorf("must contain %d-%d characters", min, max)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateLength(value, 6, 35); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, numbers and underscores")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateLength(value, 6, 35)
}

func ValidateFullName(value string) error {
	if err := ValidateLength(value, 6, 100); err != nil {
		return err
	}

	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters and spaces")
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateLength(value, 3, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("must be a valid email address")
	}
	return nil
}
