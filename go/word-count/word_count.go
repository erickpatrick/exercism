package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	cleanPhrase := strings.ToLower(phrase)
	cleanPhrase = strings.ReplaceAll(cleanPhrase, "\n", ".")
	cleanPhrase = strings.ReplaceAll(cleanPhrase, "\t", ".")
	cleanPhrase = strings.ReplaceAll(cleanPhrase, " ", ".")

	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	cleanPhrase = nonAlphanumericRegex.ReplaceAllString(cleanPhrase, " ")

	cleanPhrase = strings.ReplaceAll(cleanPhrase, " t ", "'t ")
	cleanPhrase = strings.ReplaceAll(cleanPhrase, " re ", "'re ")

	words := strings.Split(strings.TrimSpace(cleanPhrase), " ")

	freq := make(map[string]int)

	for _, word := range words {
		if value, exists := freq[word]; !exists {
			freq[word] = 1
		} else {
			freq[word] = value + 1
		}
	}

	return freq
}
