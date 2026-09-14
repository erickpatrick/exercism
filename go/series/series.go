package series

func All(n int, s string) []string {
	result := []string{}

	if n > len(s) || n < 1 {
		return result
	}

	for i := range s {
		word := ""
		if i+n > len(s) {
			word = s[i:]
		} else {
			word = s[i : i+n]
		}

		if len(word) < n {
			break
		}

		result = append(result, word)
	}

	return result
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) || n < 1 {
		return ""
	}

	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	first = UnsafeFirst(n, s)
	if first == "" {
		return first, false
	}

	return first, true
}
