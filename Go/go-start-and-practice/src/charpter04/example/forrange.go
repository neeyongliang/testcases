package main

import (
	"fmt"
)

func testArrayRange() {
	fmt.Println("test slice range:")
	for key, value := range []int{1, 2, 3, 5} {
		fmt.Printf("key:%d value:%d\n", key, value)
	}
}

func testStringRange() {
	fmt.Println("test string range:")
	var str = "hello你好"
	for key, value := range str {
		fmt.Printf("key: %d value: 0x%x\n", key, value)
	}
}

func testMapRange() {
	fmt.Println("test map range:")
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func testChannelRange() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func testFilterRange() {
	m := map[string]int{
		"go": 111,
		"zz": 222,
	}

	// 匿名变量
	for _, value := range m {
		fmt.Println(value)
	}
}

func main() {
	testArrayRange()
	testStringRange()
	testMapRange()
	testChannelRange()
	testFilterRange()
}
