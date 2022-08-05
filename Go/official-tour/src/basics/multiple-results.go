package basics

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func TestSwap(x, y string) {
	a, b := swap(x, y)
	fmt.Println(a, b)
}