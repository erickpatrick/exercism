package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	alphabet := map[rune]bool{}
	input = strings.ToLower(input)

	for _, r := range input {
		if unicode.IsLetter(r) {
			alphabet[r] = true
		}
	}

	return len(alphabet) == 26
}
