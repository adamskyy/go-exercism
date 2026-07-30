package armstrongnumbers
import "math"

func IsNumber(n int) bool {
	numberOfDigits := 0
	digits := []int{}
	original := n
	for n > 0 {
		digits = append(digits, n%10)
		numberOfDigits++
		n /= 10
	}
	sum := 0
	for _, v := range digits {
		sum += int(math.Pow(float64(v), float64(numberOfDigits)))
	}
	if sum == original {
		return true
	}
	return false
}
