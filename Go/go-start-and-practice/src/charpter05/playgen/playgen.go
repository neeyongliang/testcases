package main

import (
	"fmt"
)

func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

func main() {
	generator := playerGen("wiki")
	name, hp := generator()
	fmt.Println(name, hp)
}
