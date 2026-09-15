package largestseriesproduct

import (
	"errors"
	"regexp"
	"strings"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 1 {
		return 0, errors.New("invalid span")
	}

	re := regexp.MustCompile("[a-zA-Z]+")
	if strings.TrimSpace(re.FindString(digits)) != "" {
		return 0, errors.New("invalid digits")
	}

	// parts := All(span, digits)

	// fmt.Println(parts)

	return 0, nil
}

// copied over from go/series exercise
func All(n int, s string) []string {
	result := []string{}

	if n > len(s) || n < 1 {
		return result
	}

	for i := range s {
		word := ""
		if i+n > len(s) {
			word = s[i:]
		} else {
			word = s[i : i+n]
		}

		if len(word) < n {
			break
		}

		result = append(result, word)
	}

	return result
}
