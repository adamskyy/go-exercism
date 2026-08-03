package largestseriesproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 0 {
        return 0, errors.New("Shorter than input")
    }
    input := []rune(digits)
    
    curMax := 1
    zeroCount := 0
    for i := 0; i < span; i++ {
        if digit, ok := getDigit(input[i]); ok {
            if digit == 0 {
                zeroCount++
            } else {
                curMax *= digit 
            }
        } else {
            return 0, errors.New("Shorter than input")
        }
    }

    largestSpan := 0
    if zeroCount == 0 {
        largestSpan = curMax
    }
    for right := span; right < len(input); right++ {
        left := right - span
        
        if digit, ok := getDigit(input[left]); ok {
            if digit == 0 {
                zeroCount--
            } else {
                curMax /= digit
            }
        } else {
            return 0, errors.New("Fatal Error")
        }
        
        if digit, ok := getDigit(input[right]); ok {
			if digit == 0 {
                zeroCount++
            } else {
                curMax *= digit
            }
        } else {
            return 0, errors.New("Fatal Error")
        } 

        if curMax > largestSpan && zeroCount == 0 {
            largestSpan = curMax
        }
    }
    
	return int64(largestSpan), nil
}

func getDigit(chr rune) (digit int, ok bool) {
    if chr < '0' || chr > '9' {
    	return digit, ok
    }
    return int(chr - '0'), true
}