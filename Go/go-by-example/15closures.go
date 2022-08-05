package main

import "fmt"

func intSeq() func() int {
	fmt.Printf("intSeq")
	i := 0
	return func() int {
		fmt.Println("func-->")
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
