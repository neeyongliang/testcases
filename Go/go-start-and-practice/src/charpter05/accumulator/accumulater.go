package main

import (
	"fmt"
)

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func main() {
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())
}
