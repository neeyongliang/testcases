package main

import (
	"fmt"
)

func main() {
	seq := []string{"a", "b", "c", "d", "e"}
	index := 2
	fmt.Println(seq[:index], seq[index+1:])
	// 这里不加 ... 会报错
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}
