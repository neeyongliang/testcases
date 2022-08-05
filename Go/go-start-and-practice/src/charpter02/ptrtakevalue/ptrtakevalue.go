package main

import (
	"fmt"
)

func main() {
	var house = "Malibu Point 10880, 90265"
	ptr := &house
	fmt.Printf("ptr value: %p\n", ptr)
	value := *ptr
	fmt.Printf("value type: %T, value: %s\n", value, value)
}
