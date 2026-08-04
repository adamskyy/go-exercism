package lineup

import "fmt"

func Format(name string, number int) string {
    last := number % 10
    suffix := ""

    switch last {
        case 1: suffix = "st"
        case 2: suffix = "nd"
        case 3: suffix = "rd"
        default: suffix = "th"
    }

    
    lastTwo := number % 100
	if lastTwo == 11 || lastTwo == 12 || lastTwo == 13 {
        suffix = "th"
    }
    
    
	output := fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, suffix)
    return output
}
