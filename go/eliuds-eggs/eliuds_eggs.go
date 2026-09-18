package eliudseggs

func EggCount(displayValue int) int {
	result := 0

	for {
		if displayValue%2 == 1 {
			result += 1
		}
		displayValue /= 2
		if displayValue == 0 {
			break
		}
	}

	return result
}
