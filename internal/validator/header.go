package validator

import (
	"fmt"
	"strings"
)

func ValidateHeaderName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return fmt.Errorf("header name cannot be empty")
	}

	for _, character := range name {
		if !isValidHeaderCharacter(character) {
			return fmt.Errorf(
				"invalid character %q in header name",
				character,
			)
		}
	}

	return nil
}

func isValidHeaderCharacter(character rune) bool {
	return character >= 'a' && character <= 'z' ||
		character >= 'A' && character <= 'Z' ||
		character >= '0' && character <= '9' ||
		strings.ContainsRune("!#$%&'*+-.^_`|~", character)
}

func ValidateHeaderValue(value string) error {
	if strings.ContainsAny(value, "\r\n") {
		return fmt.Errorf(
			"header value cannot contain newline characters",
		)
	}

	return nil
}

func ValidateHeaders(headers map[string]string) error {
	for name, value := range headers {
		if err := ValidateHeaderName(name); err != nil {
			return fmt.Errorf("invalid header %q: %w", name, err)
		}

		if err := ValidateHeaderValue(value); err != nil {
			return fmt.Errorf(
				"invalid value for header %q: %w",
				name,
				err,
			)
		}
	}

	return nil
}
