package rnatranscription

func ToRNA(dna string) string {
	dnaRunes := []rune(dna)
    for i := 0; i < len(dnaRunes); i++ {
        switch cur := dnaRunes[i]; cur {
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
