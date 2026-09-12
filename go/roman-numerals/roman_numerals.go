package romannumerals

import (
	"errors"
)

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "", errors.New("number must be between 1 and 3999")
	}

	thousands := input / 1000
	hundreds := (input / 100) % 10
	tens := (input / 10) % 10
	units := input % 10

	result := ""

	for range thousands {
		result = result + "M"
	}

	if hundreds == 9 {
		result = result + "CM"
	} else if hundreds >= 5 {
		result = result + "D"

		for range hundreds - 5 {
			result = result + "C"
		}
	} else if hundreds == 4 {
		result = result + "CD"
	} else {
		for range hundreds {
			result = result + "C"
		}
	}

	if tens == 9 {
		result = result + "XC"
	} else if tens >= 5 {
		result = result + "L"

		for range tens - 5 {
			result = result + "X"
		}
	} else if tens == 4 {
		result = result + "XL"
	} else {
		for range tens {
			result = result + "X"
		}
	}

	if units == 9 {
		result = result + "IX"
	} else if units >= 5 {
		result = result + "V"

		for range units - 5 {
			result = result + "I"
		}
	} else if units == 4 {
		result = result + "IV"
	} else {
		for range units {
			result = result + "I"
		}
	}

	return result, nil
}
