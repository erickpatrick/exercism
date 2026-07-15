package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func getPhoneNumberParts(phoneNumber string) ([]string, error) {
	checker := regexp.MustCompile(`^\+?1?\s*\(?([2-9]{1}[0-9]{2})\)?[-\.\s]*([2-9]{1}[0-9]{2})[-\.\s]*([0-9]{4})[-\.\s]*$`)
	matches := checker.FindStringSubmatch(phoneNumber)

	if len(matches) < 4 {
		return matches, errors.New("invalid number")
	}

	return matches, nil
}

func Number(phoneNumber string) (string, error) {
	matches, error := getPhoneNumberParts(phoneNumber)

	if error != nil {
		return "", error
	}

	return strings.Join(matches[1:], ""), nil
}

func AreaCode(phoneNumber string) (string, error) {
	matches, error := getPhoneNumberParts(phoneNumber)

	if error != nil {
		return "", error
	}

	return matches[1], nil
}

func Format(phoneNumber string) (string, error) {
	matches, error := getPhoneNumberParts(phoneNumber)

	if error != nil {
		return "", error
	}

	return fmt.Sprintf("(%s) %s-%s", matches[1], matches[2], matches[3]), nil
}
