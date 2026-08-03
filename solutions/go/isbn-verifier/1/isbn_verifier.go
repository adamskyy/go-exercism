package isbnverifier

func IsValidISBN(isbn string) bool {
    numberOfElementsLeft := 10
    sum := 0
    number := 0
	for _, v := range isbn {
        if v == '-' {
            continue
        } else if v == 'X' && numberOfElementsLeft == 1 {
            number = 10
        } else {
            if v >= '0' && v <= '9' {
            	number = int(v - '0')
			} else {
                return false
            }
        }
        sum += (numberOfElementsLeft * number)
        numberOfElementsLeft--
    }
    if numberOfElementsLeft == 0 {
        return sum % 11 == 0
    }
    return false
}
