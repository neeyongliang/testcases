package flowcontrol

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z = (z + x / z)/2
	}
	return z
}

func ExerciseFlowControl() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(math.Sqrt(2)))
}

