package moretypes

import "fmt"

var primes = [6]int{2, 3, 5, 7, 11, 13}

func TestArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(primes)
}

func TestArraySlice() {
	var s []int = primes[1:4]
	fmt.Println(s)
}
