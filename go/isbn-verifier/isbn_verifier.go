package isbnverifier

import (
	"fmt"
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
	i := 0

	for {
	}

	fmt.Println(cleanISBN)

	return sum%11 == 0
}
