package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile("<[~*=-]*>")
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	counter := 0

	for _, line := range lines {
		result := re.FindAllString(line, -1)
		counter += len(result)
	}

	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile("User +([A-Z]{1}[a-z]+[0-9]+)")

	for i, line := range lines {
		result := re.FindAllStringSubmatch(line, -1)
		if len(result) == 1 {
			lines[i] = fmt.Sprintf("[USR] %s %s", result[0][1], line)
		}
	}

	return lines
}
