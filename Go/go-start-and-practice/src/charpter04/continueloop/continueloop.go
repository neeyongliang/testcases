package main

import (
	"fmt"
)

func main() {
OuterLoop2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop2
			}
		}
	}
}

// Outputs:
// 0 2
// 1 2
