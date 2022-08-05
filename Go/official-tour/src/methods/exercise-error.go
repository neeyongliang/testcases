package methods

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := x
	for i :=0; i < 10; i++ {
		z = (z + x / z) /2
	}
	return z, nil
}

func ExerciseError() {
	if value, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	if value, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}
