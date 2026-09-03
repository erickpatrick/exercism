package rnatranscription

import "strings"

func ToRNA(dna string) string {
	result := ""
	for _, nucleotide := range strings.ToUpper(dna) {
		switch nucleotide {
		case 'G':
			result = result + "C"
		case 'C':
			result = result + "G"
		case 'T':
			result = result + "A"
		case 'A':
			result = result + "U"
		default:
			continue
		}
	}

	return result
}
