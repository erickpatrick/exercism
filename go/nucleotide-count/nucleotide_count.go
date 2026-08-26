package nucleotidecount

import "errors"

type (
	Histogram map[rune]int
	DNA       string
)

func (d DNA) Counts() (Histogram, error) {
	var h Histogram = map[rune]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, nucleotide := range d {
		switch nucleotide {
		case 'A', 'C', 'G', 'T':
			h[nucleotide]++
		default:
			return h, errors.New("error")
		}
	}

	return h, nil
}
