package rnatranscription

func ToRNA(dna string) string {
	dnaRunes := make([]rune, len(dna))
    for i := 0; i < len(dna); i++ {
        switch cur := dna[i]; cur {
        case 'G': 
            dnaRunes[i] = 'C'
        case 'C': 
            dnaRunes[i] = 'G'
        case 'T': 
            dnaRunes[i] = 'A'
        case 'A': 
            dnaRunes[i] = 'U'
		default:
        }
    }
    return string(dnaRunes)
}
