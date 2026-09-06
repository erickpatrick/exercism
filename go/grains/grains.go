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
	panic("Please implement the Total function")
}
