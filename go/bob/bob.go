// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// fmt.Println("`", remark, "`")
	remark = strings.TrimSpace(remark)
	// fmt.Println("`", remark, "`")
	// fmt.Println("---")

	isEmpty := regexp.MustCompile(`^\s+$`)

	if isEmpty.MatchString(remark) || remark == "" {
		return "Fine. Be that way!"
	}

	isQuestion := remark[len(remark)-1] == '?'
	isYelling := regexp.MustCompile(`([[:upper:]]|[0-9])\b{2,}`)
	isNormal := regexp.MustCompile(`([[:lower:]]|[0-9])+`)

	if isQuestion {
		if isYelling.MatchString(remark) && !isNormal.MatchString(remark) {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}

	if isYelling.MatchString(remark) && !isNormal.MatchString(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
