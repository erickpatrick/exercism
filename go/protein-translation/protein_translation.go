package proteintranslation

import "errors"

var (
	ErrStop        = errors.New("STOP")
	ErrInvalidBase = errors.New("invalid RNA sequence")
)

func codonToAminoacid(codon string) string {
	codons := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	return codons[codon]
}

func FromRNA(rna string) ([]string, error) {
	var result []string

	if len(rna)%3 != 0 && len(rna) < 6 {
		return result, ErrInvalidBase
	}

	iterations := len(rna) / 3
	for i := 1; i <= iterations; i += 1 {
		aminoacid, error := FromCodon(rna[3*(i-1) : 3*i])

		if error == ErrStop {
			break
		}

		if error == ErrInvalidBase {
			return result, ErrInvalidBase
		}

		result = append(result, aminoacid)
	}

	return result, nil
}

func FromCodon(codon string) (string, error) {
	aminoacid := codonToAminoacid(codon)

	if aminoacid == "" {
		return aminoacid, ErrInvalidBase
	}

	if aminoacid == "STOP" {
		return aminoacid, ErrStop
	}

	return aminoacid, nil
}
