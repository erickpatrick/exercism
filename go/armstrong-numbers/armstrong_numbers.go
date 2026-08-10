package armstrongnumbers

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	nAsString := strconv.Itoa(n)
	nLength := len(nAsString)
	total := 0

	for _, digit := range nAsString {
		i, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
		total += int(math.Pow(float64(i), float64(nLength)))
	}

	return total == n
}
