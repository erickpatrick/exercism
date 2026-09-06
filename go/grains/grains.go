package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("")
	}
	return uint64(math.Pow(float64(2), float64(number-1))), nil
}

func Total() uint64 {
	var total uint64

	for i := 1; i < 65; i++ {
		value, err := Square(i)
		if err != nil {
			total += 0
		}
		total += value
	}

	return total
}
