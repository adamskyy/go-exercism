package phonenumber

import (
    "errors"
    "fmt"
)



func Number(phoneNumber string) (string, error) {
	var cleaned []byte

    for i := 0; i < len(phoneNumber); i++ {
        c := phoneNumber[i]
        switch {
        case c >= '0' && c <= '9':
            cleaned = append(cleaned, c)
        case c == ' ' || c == '-' || c == '.' || c == '+' || c == '(' || c == ')':
			continue    
        default:
            return "", errors.New("invalid character")
        }
    }

    if len(cleaned) == 11 && cleaned[0] == '1' {
        cleaned = cleaned[1:]
    } else if len(cleaned) != 10 {
        return "", errors.New("invalid length")
    }

    // Additional checks
    if (cleaned[0] != '1') && (cleaned[3] != '1') && (cleaned[0] != '0') && (cleaned[3] != '0') {
        return string(cleaned), nil
    }
    return "", errors.New("invalid character")
}

func AreaCode(phoneNumber string) (string, error) {
	cleanedNumber, err := Number(phoneNumber)
    if err != nil {
        return "", errors.New("invalid number")
    }
    return cleanedNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanedNumber, err := Number(phoneNumber)
    if err != nil {
        return "", errors.New("invalid number")
    }
    return fmt.Sprintf("(%s) %s-%s", cleanedNumber[:3], cleanedNumber[3:6], cleanedNumber[6:]), nil
}
