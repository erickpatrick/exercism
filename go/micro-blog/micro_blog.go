package microblog

import "fmt"

func Truncate(phrase string) (result string) {
	counter := 1

	for _, r := range phrase {
		result += fmt.Sprintf("%c", r)
		counter++
		if counter > 5 {
			break
		}
	}

	return result
}
