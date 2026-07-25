package isogram

import "strings"

func IsIsogram(word string) bool {
    wordLower := strings.ToLower(word)
	seen := map[rune]struct{}{}
    for _, char := range wordLower {
        if char < 'a' || char > 'z' {
            continue
        }
        if _, exists := seen[char]; exists {
            return false
        }
        seen[char] = struct{}{}
    }
    return true
}
