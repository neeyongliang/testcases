package main

import "fmt"

type class struct {
}

func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

func funcDo(v int) {
	fmt.Println("call function do:", v)
}

func main() {
	// 声明一个函数回调
	var delegate func(int)

	c := new(class)
	delegate = c.Do
	delegate(100)

	delegate = funcDo
	delegate(100)
}
