// Package for generation of acronyms
package acronym

import "unicode"

// Function abbreviate takes as input a string, and it returns an acronym
func Abbreviate(s string) string {
	output := []rune{}
    takeNext := true
    for _, char := range s {
        if char == '-' || char == ' ' {
            takeNext = true
            continue
        }

        if takeNext && unicode.IsLetter(char) {
            output = append(output, unicode.ToUpper(char))
            takeNext = false
        }
    }
    return string(output)
}
