// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// fmt.Println("`", remark, "`")
	remark = strings.TrimSpace(remark)
	// fmt.Println("`", remark, "`")
	// fmt.Println("---")

	if remark == "" {
		return "Fine. Be that way!"
	}

	hasLower := false
	hasUpper := false
	hasQuestionMark := remark[len(remark)-1] == '?'

	for _, r := range remark {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				hasUpper = true
			}

			if unicode.IsLower(r) {
				hasLower = true
			}
		}
	}

	if hasQuestionMark {
		if hasUpper && !hasLower {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}

	if hasUpper && !hasLower {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
