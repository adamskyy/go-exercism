package microblog

import (
	"unicode/utf8"
)

func Truncate(phrase string) string {
    count := 0
	bytePosition := 0
	for len(phrase[bytePosition:]) > 0 && count < 5 {
		_, runeSize := utf8.DecodeRuneInString(phrase[bytePosition:])
		bytePosition += runeSize
		count++
	}
	return phrase[:bytePosition]
}
