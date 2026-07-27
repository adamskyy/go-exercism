package luhn

import (
    "strings"
)

func Valid(id string) bool {
    inputWithoutSpaces := strings.ReplaceAll(id, " ", "")
	if !ValidateInput(inputWithoutSpaces) {
        return false
    }
    return validateCheckSum(inputWithoutSpaces)
}

func ValidateInput(id string) bool {
    if len(id) <= 1 {
        return false
    }
    for _, char := range id {
        if char < '0' || char > '9' {
            return false
        }
    }
    return true
}

func validateCheckSum(id string) bool {
    sum := 0
    doubleDigit := false
    for i := len(id) - 1; i >= 0; i-- {
        curNumber := int(id[i] - '0')
		if doubleDigit {
            curNumber *= 2
            if curNumber >= 9 {
            	curNumber -= 9
       		}
        }
        sum += curNumber
    	doubleDigit = !doubleDigit
    }
    return sum % 10 == 0
}