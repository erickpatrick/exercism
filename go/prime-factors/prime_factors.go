package primefactors

func Factors(n int64) []int64 {
	result := []int64{}

	for i := int64(2); i <= n; {
		if n%i == 0 {
			result = append(result, i)
			n = n / i
			continue
		}

		i++
	}

	return result
}
