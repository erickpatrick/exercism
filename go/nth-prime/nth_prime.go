package nthprime

import (
	"errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("number cannot be smaller or equal to 0")
	}
	primes := []int{}
	counter := 2

	for {
		if isItPrime(counter) {
			primes = append(primes, counter)
		}

		if len(primes) >= n || counter > 50 {
			break
		}

		counter++
	}

	return primes[n-1], nil
}

func isItPrime(n int) bool {
	divisors := 0
	for counter := 1; counter <= n; counter++ {
		if n%counter == 0 {
			divisors++
		}

		if divisors > 2 {
			return false
		}
	}

	return divisors == 2
}
