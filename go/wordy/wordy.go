package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	if !strings.HasPrefix(question, "What is ") || !strings.HasSuffix(question, "?") {
		return 0, false
	}

	question = strings.Replace(question, "What is ", "", 1)
	question = strings.Replace(question, "?", "", 1)
	parts := strings.Split(question, " ")

	switch len(parts) {
	case 1:
		return naturalNumber(parts[0])
	case 3, 4:
		return doOperation(parts)
	}

	// group1 = first number
	// group3 =

	return 0, false
}

func naturalNumber(num string) (int, bool) {
	i, err := strconv.Atoi(num)
	if err != nil {
		return 0, false
	}

	return i, true
}

func doOperation(parts []string) (int, bool) {
	fmt.Println(parts)
	switch parts[1] {
	case "plus":
		return sum(parts[0], parts[2])
	case "minus":
		return subtract(parts[0], parts[2])
	case "multiplied":
		return multiply(parts[0], parts[3])
	case "divided":
		return divide(parts[0], parts[3])
	default:
		return 0, false
	}
}

func sum(num1 string, num2 string) (int, bool) {
	val1, exists := naturalNumber(num1)
	if !exists {
		return 0, false
	}

	val2, exists := naturalNumber(num2)
	if !exists {
		return 0, false
	}
	return val1 + val2, true
}

func subtract(num1 string, num2 string) (int, bool) {
	val1, exists := naturalNumber(num1)
	if !exists {
		return 0, false
	}

	val2, exists := naturalNumber(num2)
	if !exists {
		return 0, false
	}
	return val1 - val2, true
}

func multiply(num1 string, num2 string) (int, bool) {
	val1, exists := naturalNumber(num1)
	if !exists {
		return 0, false
	}

	val2, exists := naturalNumber(num2)
	if !exists {
		return 0, false
	}
	return val1 * val2, true
}

func divide(num1 string, num2 string) (int, bool) {
	val1, exists := naturalNumber(num1)
	if !exists {
		return 0, false
	}

	val2, exists := naturalNumber(num2)
	if !exists {
		return 0, false
	}
	return val1 / val2, true
}
