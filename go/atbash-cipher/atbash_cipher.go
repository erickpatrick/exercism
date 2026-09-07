package atbashcipher

import (
	"strings"
	"unicode"
)

const (
	startLetter = 97
	finalLetter = 122
)

func Atbash(s string) string {
	final := ""
	// convert letters
	for _, letter := range strings.ToLower(s) {
		if unicode.IsNumber(letter) {
			final = final + string(letter)
		}

		if unicode.IsLetter(letter) && (letter >= startLetter || letter <= finalLetter) {
			final = final + string(rune(finalLetter-(letter-startLetter)))
		}

		// split in groups of 5
		if len(final)%6 == 5 {
			final = final + " "
		}
	}

	return strings.TrimSuffix(final, " ")
}
