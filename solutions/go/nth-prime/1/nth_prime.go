package nthprime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Cannot be computed")
	}
	if n == 1 {
		return 2, nil
	}
	if n == 2 {
		return 3, nil
	}
	primeDedected := 1
	i := 1
	for primeDedected != n {
        i += 2
		if isPrime(i) {
			primeDedected++
		}
	}
	return i, nil
}

func isPrime(n int) bool {
	prime := true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}

	return prime
}
