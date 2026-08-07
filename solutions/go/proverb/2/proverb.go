
// Package proverb should have a package comment that summarizes what it's about.
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
    var lines []string
    if len(rhyme) == 0 {
        return lines
    }
    
    for i := 0; i < len(rhyme) - 1; i++ {
        newLine := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
        lines = append(lines, newLine)
    }
    lastLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
    
	return append(lines, lastLine)
}
