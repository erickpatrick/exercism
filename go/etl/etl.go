package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)

	for value, letters := range in {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = value
		}
	}

	return result
}
