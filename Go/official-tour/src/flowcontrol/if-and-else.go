package flowcontrol

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("%g >= %g\n", v, lim)
	}
	return lim
}

func TestIfAndElse() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
