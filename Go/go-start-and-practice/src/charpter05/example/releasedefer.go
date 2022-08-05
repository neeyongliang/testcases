package main

import (
	"fmt"
	"os"
)

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}

	info, err := f.Stat()

	if err != nil {
		f.Close()
		return 0
	}

	size := info.Size()
	return size
}

func fileSize2(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%s not found\n", filename)
		return 0
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		fmt.Printf("%s get state error\n", filename)
		return 0
	}

	size := info.Size()
	return size
}

func main() {
	size := fileSize2("/Users/wiki/.zshrc")
	fmt.Println(size)
}
