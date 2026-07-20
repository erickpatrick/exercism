package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string
	subjectRunes := []rune(strings.ToLower(subject))
	subjectLower := strings.ToLower(subject)

	for _, candidate := range candidates {
		candidateRunes := []rune(strings.ToLower(candidate))
		candidateLower := strings.ToLower(candidate)

		// orders runes to properly compare them
		slices.Sort(subjectRunes)
		slices.Sort(candidateRunes)

		if slices.Compare(subjectRunes, candidateRunes) == 0 &&
			subjectLower != candidateLower {
			result = append(result, candidate)
		}
	}

	return result
}
