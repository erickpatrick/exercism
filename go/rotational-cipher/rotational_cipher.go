package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	result := []rune{}

	for _, r := range plain {
		if shiftKey == 0 {
			result = append(result, r)
			continue
		}

		if shiftKey > 26 {
			shiftKey -= 26
		}

		if unicode.IsLetter(r) {
			if unicode.IsLower(r) {
				if unicode.IsLower(rune(int(r) + shiftKey)) {
					result = append(result, rune(int(r)+shiftKey))
				} else {
					// 123 so it can start from a
					result = append(result, rune(97+(int(r)+shiftKey-123)))
				}
			} else {
				if unicode.IsUpper(rune(int(r) + shiftKey)) {
					result = append(result, rune(int(r)+shiftKey))
				} else {
					// 91 so it can start from A
					result = append(result, rune(65+(int(r)+shiftKey-91)))
				}
			}
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}
