package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	alphabet := map[rune]bool{}
	input = strings.ToLower(input)

	if len(input) == 0 {
		return false
	}

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}

		alphabet[r] = true
	}

	return len(alphabet) == 26
}
