// Package leap implement a function to identify leap years

package leap

// IsLeapYear returns a flag whether given year is a leap year
func IsLeapYear(year int) bool {
	if year % 4 == 0 {
        if year % 100 == 0 {
            if year % 400 == 0 {
                return true
            }
            return false
        }
        return true
    }
    return false
}
