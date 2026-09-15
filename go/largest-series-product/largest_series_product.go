package largestseriesproduct

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
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

	products := []int64{}
	sequencies := All(span, digits)

	for _, sequence := range sequencies {
		product := int64(1)
		for _, value := range sequence {
			val, _ := strconv.Atoi(string(value))
			product = product * int64(val)
		}
		products = append(products, product)
	}

	slices.Sort(products)

	return products[len(products)-1], nil
}

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
