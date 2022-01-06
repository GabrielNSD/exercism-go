package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	result, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}

	if err != nil {
		if err == ErrScaleMalfunction && result > 0 {
			return (result * 2) / float64(cows), nil
		} else if err == ErrScaleMalfunction && result < 0 {
			return 0, errors.New("negative fodder")
		}
		return 0, err
	} else {
		if result < 0 {
			return 0, errors.New("negative fodder")
		}
		return result / float64(cows), err
	}
}
