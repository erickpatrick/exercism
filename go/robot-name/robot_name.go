package robotname

import (
	"errors"
	"math/rand"
	"strings"
)

type Robot struct {
	name string
}

const (
	maxRobotNames = 26 * 26 * 10 * 10 * 10
	// to simplify code and simply use  randInt63() instead of making more
	// calculations
	letters = "abcdefghijklmnopqrstuwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
)

var robotNames = map[string]bool{}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := generateRobotName()
		if err != nil {
			return "", errors.New("no new name possible")
		}
		r.name = name
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	name, _ := generateRobotName()
	r.name = name
}

func generateRobotName() (string, error) {
	if len(robotNames) == maxRobotNames {
		return "", errors.New("reached maximum names possible")
	}

	var robotName string
	next := true

	// uses for loop as otheriwse the recursive version would have ended up
	// into stackoverflow
	for next {
		prefix := make([]byte, 2)
		for i := range prefix {
			prefix[i] = letters[rand.Int63()%int64(len(letters))]
		}

		appendix := make([]byte, 3)
		for i := range appendix {
			appendix[i] = numbers[rand.Int63()%int64(len(numbers))]
		}

		robotName = strings.ToUpper(string(prefix) + string(appendix))

		if _, ok := robotNames[robotName]; !ok {
			next = false
		}
	}

	robotNames[robotName] = true

	return robotName, nil
}
