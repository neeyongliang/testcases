package main

import (
	"fmt"
)

// 宕机
func main() {
	defer fmt.Println("Todo things after panic")
	defer fmt.Println("Todo things after panic 2")
	panic("panic")
}
