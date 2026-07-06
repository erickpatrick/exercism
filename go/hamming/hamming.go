package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands with different sizes")
	}

	if len(a) == 0 {
		return 0, nil
	}

	distance := 0
	i := 0
	for i < len(a) {
		if a[i] != b[i] {
			distance += 1
		}
		i += 1
	}

	return distance, nil
}
