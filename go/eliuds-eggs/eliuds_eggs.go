package eliudseggs

func EggCount(displayValue int) int {
	result := 0

	for displayValue > 0 {
		if displayValue%2 == 1 {
			result += 1
		}
		displayValue /= 2
	}

	return result
}
