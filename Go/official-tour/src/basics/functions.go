package basics

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func TestAdd(x int, y int) {
	fmt.Println(add(x, y))
}

func TestAdd2(x, y int) {
	fmt.Println(add(x, y))
}
