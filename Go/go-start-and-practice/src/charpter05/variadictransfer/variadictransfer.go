package main

import (
	"fmt"
)

func rawPrintf(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

func print(slist ...interface{}) {
	rawPrintf(slist...)
}

func main() {
	println(1, 2, 3)
}
