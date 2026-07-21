package sorting

import (
	"fmt"
	"strconv"
)
// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float32(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    fnbValid, ok := fnb.(FancyNumber)
    if !ok {
        return 0
    }
	val, err := strconv.Atoi(fnbValid.Value())
    if err != nil {
        return 0
    }
    return val
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    fnbValid, ok := fnb.(FancyNumber)
    if ok {
        val, _ := strconv.Atoi(fnbValid.Value())
        return fmt.Sprintf("This is a fancy box containing the number %.1f", float32(val))
    }
    return fmt.Sprintf("This is a fancy box containing the number 0.0")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
    var output string
	switch v := i.(type) {
        case int:
            output = DescribeNumber(float64(v))
    	case float64:
        	output = DescribeNumber(v)
        case NumberBox:
            output = DescribeNumberBox(v)
        case FancyNumberBox:
            output = DescribeFancyNumberBox(v)
        default:
            output = "Return to sender"
    }
    return output
}
