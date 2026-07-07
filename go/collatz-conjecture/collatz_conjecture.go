package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("number must be positive")
	}

	steps := 0
	currentNumber := n
	for {
		if currentNumber == 1 {
			return steps, nil
		}

		if currentNumber%2 == 0 {
			currentNumber = currentNumber / 2
			steps++
			continue
		}

		if currentNumber%2 != 0 {
			currentNumber = (currentNumber * 3) + 1
			steps++
			continue
		}
	}
}
