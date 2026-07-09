package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	if len(id) <= 1 {
		return false
	}

	checker := regexp.MustCompile(`^[\d\s]+$`)
	isNumericID := checker.MatchString(id)
	if !isNumericID {
		return false
	}

	cleanID := strings.ReplaceAll(id, " ", "")
	sum := 0

	if len(cleanID) <= 1 {
		return false
	}

	runes := []rune(cleanID)
	maxLength := len(runes) - 1
	for i := range runes {
		val, err := strconv.Atoi(string(runes[maxLength-i]))
		if err != nil {
			continue
		}

		if i%2 == 0 {
			sum += val
			continue
		}

		doubleValue := val * 2
		if doubleValue > 9 {
			sum += doubleValue - 9
		} else {
			sum += doubleValue
		}
	}

	return sum%10 == 0
}
