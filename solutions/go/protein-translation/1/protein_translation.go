package proteintranslation

import "errors" 
var mapping = map[string]string{
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

var ErrStop = errors.New("stop amino")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
    var outputs []string
	lenRna := len(rna)
    for i := 0; i + 3 <= lenRna; i += 3 {
        aminoAcid, err := FromCodon(rna[i: i+3])
        if err == ErrInvalidBase {
    		return outputs, err
    	} else if err == ErrStop {
			return outputs, nil
        }
        outputs = append(outputs, aminoAcid)
    }
    if lenRna % 3 != 0 {
        return outputs, ErrInvalidBase
    }
    return outputs, nil
}

func FromCodon(codon string) (string, error) {
    elem, ok := mapping[codon]
    if !ok {
    	return "", ErrInvalidBase
    } else if elem == "STOP" {
        return "", ErrStop
    }
    return elem, nil
}