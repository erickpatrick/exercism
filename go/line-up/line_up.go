package lineup

import "fmt"

func Format(name string, number int) string {
	writtenNumber := fmt.Sprintf("%d", number)
	rest := number % 100

	if rest > 13 {
		rest = rest % 10
	}

	switch true {
	case number == 11 || number == 12 || number == 13:
		writtenNumber += "th"
	case rest == 3:
		writtenNumber += "rd"
	case rest == 2:
		writtenNumber += "nd"
	case rest == 1:
		writtenNumber += "st"
	default:
		writtenNumber += "th"
	}

	return fmt.Sprintf("%s, you are the %s customer we serve today. Thank you!", name, writtenNumber)
}
