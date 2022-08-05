package main

import (
	"fmt"
)

func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
}

// 匿名结构体每次都需要重新定义，造成大量重复代码，很少使用

func main() {
	msg := &struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}

	printMsgType(msg)
}
