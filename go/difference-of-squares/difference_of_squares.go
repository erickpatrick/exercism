package differenceofsquares

func SquareOfSum(n int) int {
	if n == 0 {
		return 0
	}

	sum := (1 + n) * n / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
