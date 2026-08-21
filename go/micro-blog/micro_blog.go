package microblog

func Truncate(phrase string) (result string) {
	r := []rune(phrase)

	if len(r) < 5 {
		return string(r)
	}

	return string(r[:5])
}
