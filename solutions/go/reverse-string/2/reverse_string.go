package reversestring

func Reverse(input string) string {
    runes := []rune(input)
	left, right := 0, len(runes)-1

    for left < right {
    
        runes[left], runes[right] = runes[right], runes[left]
    
        left++
    
        right--
    
    }
    return string(runes)
}
