package main

import (
	"fmt"
)

func test_ascii() {
	theme := "狙击枪start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c,  %d\n", theme[i], theme[i])
	}
}

func test_unicode() {
	theme := "狙击枪start"
	for _, s := range theme {
		fmt.Printf("Unicode: %c  %d\n", s, s)
	}
}

func main() {
	test_ascii()
	test_unicode()
}
