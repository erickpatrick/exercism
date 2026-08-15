package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)

	// checks for word with possible word'word combination
	regex, _ := regexp.Compile(`[\w]+(?:'[\w]+)*`)
	words := regex.FindAllString(phrase, -1)
	freq := make(Frequency)

	for _, word := range words {
		freq[word]++
	}
	return freq
}
