package main

import (
	"fmt"
)

func main() {
	fmt.Println("defer begin")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("defer end")
}

// outputs:
// defer begin
// defer end
// 3
// 2
// 1
