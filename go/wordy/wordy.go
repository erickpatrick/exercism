package wordy

import (
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
		return doSimpleOperation(parts)
	case 5, 6, 7:
		return doMultipleOperation(parts)
	}

	return 0, false
}

func naturalNumber(num string) (int, bool) {
	i, err := strconv.Atoi(num)
	if err != nil {
		return 0, false
	}

	return i, true
}

func doSimpleOperation(parts []string) (int, bool) {
	if len(parts) == 4 && parts[2] != "by" {
		return 0, false
	}

	val1, exists := naturalNumber(parts[0])
	if !exists {
		return 0, false
	}

	val2, exists := naturalNumber(parts[len(parts)-1])

	if !exists {
		return 0, false
	}

	switch parts[1] {
	case "plus":
		return val1 + val2, true
	case "minus":
		return val1 - val2, true
	case "multiplied":
		return val1 * val2, true
	case "divided":
		return val1 / val2, true
	default:
		return 0, false
	}
}

func doMultipleOperation(parts []string) (int, bool) {
	limit := 3
	if parts[2] == "by" {
		limit = 4
	}
	result1, worked := doSimpleOperation(parts[0:limit])
	if !worked {
		return 0, false
	}

	newParts := []string{}
	newParts = append(newParts, strconv.Itoa(result1))
	newParts = append(newParts, parts[limit:]...)

	return doSimpleOperation(newParts)
}
