package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	// this is like fizzbuzz, but instead of writting fizzbuzz, we sum the unique
	// values which are divisible by either number in `divisors` up to `limit`

	points := map[int]int{}
	for i := range limit {
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			if i%divisor == 0 {
				points[i] = i
			}
		}
	}

	result := 0
	for _, point := range points {
		result += point
	}

	return result
}
