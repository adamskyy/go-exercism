package largestseriesproduct

import "errors"

var (

	errInvalidSpan  = errors.New("span must not be negative or greater than input length")
	errInvalidDigit = errors.New("input must contain only digits")

)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span < 0 || span > len(digits) {
		return 0, errInvalidSpan
	}
	if span == 0 {
		return 1, nil
	}
	var product int64 = 1
	zeroCount := 0
	for i := 0; i < span; i++ {
		digit, err := digitValue(digits[i])
		if err != nil {
			return 0, err
		}
		if digit == 0 {
			zeroCount++
		} else {
			product *= digit
		}
	}
	var largest int64
	if zeroCount == 0 {
		largest = product
	}
	for right := span; right < len(digits); right++ {
		left := right - span
		outgoing, err := digitValue(digits[left])
		if err != nil {
			return 0, err
		}
		if outgoing == 0 {
			zeroCount--
		} else {
			product /= outgoing
		}
		incoming, err := digitValue(digits[right])
		if err != nil {
			return 0, err
		}
		if incoming == 0 {
			zeroCount++
		} else {
			product *= incoming
		}
		if zeroCount == 0 && product > largest {
			largest = product
		}
	}
	return largest, nil

}

func digitValue(char byte) (int64, error) {

	if char < '0' || char > '9' {
		return 0, errInvalidDigit
	}
	return int64(char - '0'), nil

}