package isbnverifier

import (
	"regexp"
	"strings"
)

func IsValidISBN(isbn string) bool {
	cleanISBN := strings.ReplaceAll(isbn, "-", "")
	if len(cleanISBN) == 0 {
		return false
	}

	checker := regexp.MustCompile(`^\d{9}[\dX]{1}$`)
	if doesISBNContainValidChars := checker.MatchString(cleanISBN); !doesISBNContainValidChars {
		return false
	}

	sum := 0
	multiplier := 10

	for i := range 10 {
		switch i {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8:
			sum += int(cleanISBN[i]-'0') * multiplier
			multiplier--
		case 9:
			if rune(cleanISBN[i]) == 'X' {
				sum += 10
			} else {
				sum += int(cleanISBN[i] - '0')
			}
		}
	}

	return sum%11 == 0
}
