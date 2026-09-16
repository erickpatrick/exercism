package sieve

func Sieve(limit int) []int {
	nums := []int{}
	for i := 2; i <= limit; i++ {
		add := true
		for j := i - 1; j >= 2; j-- {
			if i%j == 0 {
				add = false
				break
			}
		}

		if add {
			nums = append(nums, i)
		}
	}

	return nums
}
