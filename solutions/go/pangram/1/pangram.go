package pangram

import (
    "unicode"
)

func IsPangram(input string) bool {
	hashmap := map[rune]struct{}{}
    
	for _, char := range input {
    	lowerCase := unicode.ToLower(char)
        if lowerCase >= 'a' && lowerCase <= 'z' {
            hashmap[lowerCase] = struct{}{}
        }
    }
    return len(hashmap) == 26
}
