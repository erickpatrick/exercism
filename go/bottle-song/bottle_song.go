package bottlesong

import (
	"fmt"
	"strings"
)

func pluralize(value int) string {
	switch value {
	case 0:
		return "no green bottles"
	case 1:
		return "One green bottle"
	case 2:
		return "Two green bottles"
	case 3:
		return "Three green bottles"
	case 4:
		return "Four green bottles"
	case 5:
		return "Five green bottles"
	case 6:
		return "Six green bottles"
	case 7:
		return "Seven green bottles"
	case 8:
		return "Eight green bottles"
	case 9:
		return "Nine green bottles"
	case 10:
		return "Ten green bottles"
	default:
		return ""
	}
}

func Recite(startBottles, takeDown int) []string {
	result := []string{}
	final := startBottles - takeDown
	for i := startBottles; i > final; i-- {
		current := pluralize(i)
		next := strings.ToLower(pluralize(i - 1))

		result = append(result, fmt.Sprintf("%s hanging on the wall,", current))
		result = append(result, fmt.Sprintf("%s hanging on the wall,", current))
		result = append(result, "And if one green bottle should accidentally fall,")
		result = append(result, fmt.Sprintf("There'll be %s hanging on the wall.", next))

		takeDown--
		if takeDown > 0 {
			result = append(result, "")
		}
	}

	return result
}
