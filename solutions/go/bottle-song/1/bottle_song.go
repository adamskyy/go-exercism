package bottlesong

import (
	"fmt"
	"strings"
)

var store = map[int]string{
	0:  "No",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func Recite(startBottles, takeDown int) []string {
	var output []string
	for takeDown > 0 {
		output = append(output, GenerateOneParagraph(startBottles)...)
		startBottles--
		takeDown--
        if takeDown > 0 {
            output = append(output, "")
        }
	}
	return output
}

func GenerateOneParagraph(bottles int) []string {
    fsuffix := "s"
    if bottles == 1 {
        fsuffix = ""
    }
	firstLines := fmt.Sprintf("%s green bottle%s hanging on the wall,", store[bottles], fsuffix)
	thirdLine := "And if one green bottle should accidentally fall,"
	left := bottles - 1
	suffix := "s"
	if left == 1 {
		suffix = ""
	}
	fourthLine := fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", strings.ToLower(store[left]), suffix)
	return []string{firstLines, firstLines, thirdLine, fourthLine}
}