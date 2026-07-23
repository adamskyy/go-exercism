package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
    amount, err := fodderCalculator.FodderAmount(numberOfCows)
    if err != nil {
        return 0.0, err
    }
    factor, errF := fodderCalculator.FatteningFactor()
    if errF != nil {
        return 0.0, errF
    }
    return amount * factor / float64(numberOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
    if numberOfCows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    }
    output, err := DivideFood(fodderCalculator, numberOfCows)
    if err != nil {
        return 0.0, err
    }
    return output, nil
}

type InvalidCowsError struct {
  numberOfCows int
  message string
}

func (e *InvalidCowsError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOfCows int) error {
    if numberOfCows < 0 {
        return &InvalidCowsError{numberOfCows, "there are no negative cows"}
    } else if numberOfCows == 0 {
        return &InvalidCowsError{numberOfCows, "no cows don't need food"}
    }
    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
