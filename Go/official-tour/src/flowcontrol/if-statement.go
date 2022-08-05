package flowcontrol

import (
	"fmt"
	"math"
)

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func TestIfStatement() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
