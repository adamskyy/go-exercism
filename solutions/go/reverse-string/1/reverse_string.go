package reversestring

func Reverse(input string) string {
	runes := []rune(input)
    out := ""
    for i := len(runes) - 1; i >= 0; i-- {
        out += string(runes[i])
    }
    return out
}
