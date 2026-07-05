package raindrops

import "strconv"

func Convert(number int) string {
	result := ""

	switch {
	case number%105 == 0:
		result = "PlingPlangPlong"
	case number%15 == 0:
		result = "PlingPlang"
	case number%21 == 0:
		result = "PlingPlong"
	case number%35 == 0:
		result = "PlangPlong"
	case number%3 == 0:
		result = "Pling"
	case number%5 == 0:
		result = "Plang"
	case number%7 == 0:
		result = "Plong"
	default:
		result = strconv.Itoa(number)
	}
	return result
}
