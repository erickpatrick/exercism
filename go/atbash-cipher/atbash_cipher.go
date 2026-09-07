package atbashcipher

import (
	"slices"
	"strings"
	"unicode"
)

const (
	startLetter = 97
	finalLetter = 122
)

func Atbash(s string) string {
	converted := []string{}
	// convert letters
	for _, letter := range strings.ToLower(s) {
		if unicode.IsNumber(letter) {
			converted = append(converted, string(letter))
		}

		if unicode.IsLetter(letter) && (letter >= startLetter || letter <= finalLetter) {
			converted = append(converted, string(rune(finalLetter-(letter-startLetter))))
		}
	}

	final := []string{}
	for chunk := range slices.Chunk(converted, 5) {
		final = append(final, strings.Join(chunk, ""))
	}

	return strings.Join(final, " ")
}
