package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "-", " ")
	parts := strings.Split(s, " ")
	result := ""

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		for _, c := range part {
			if unicode.IsLetter(c) {
				result = result + string(c)
				break
			}
		}
	}

	return result
}
