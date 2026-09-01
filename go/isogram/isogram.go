package isogram

import "strings"

func IsIsogram(word string) bool {
	entries := map[rune]bool{}
	word = strings.ToLower(word)

	for _, r := range word {
		if _, exists := entries[r]; exists && (r != '-' && r != ' ') {
			return false
		}
		entries[r] = true
	}

	return true
}
