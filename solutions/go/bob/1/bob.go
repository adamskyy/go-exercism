// Package bob classify questions.
package bob

import (
	"regexp"
	"strings"
)

var containsLetter = regexp.MustCompile(`[A-Za-z]`)

// Function Hey returns a questions to answer
func Hey(remark string) string {
	message := strings.TrimSpace(remark)

	if isSilent(message) {
		return "Fine. Be that way!"
	}

	yelling := isYelling(message)
	question := isQuestion(message)

	switch {
	case yelling && question:
		return "Calm down, I know what I'm doing!"
	case question:
		return "Sure."
	case yelling:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isSilent(message string) bool {
	return message == ""
}

func isQuestion(message string) bool {
	return strings.HasSuffix(message, "?")
}
func isYelling(message string) bool {
	// Numbers and punctuation alone should not count as yelling.
	if !containsLetter.MatchString(message) {

		return false

	}

	return message == strings.ToUpper(message)

}
