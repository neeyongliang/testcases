package main

import "fmt"

// 本函数测试入口参数和返回值的情况
// go run -gcflags "-m -l" xxx.go
func dummy(b int) int {
	// 声明一个c赋值进入函数

	var c int
	c = b

	return c
}

// 空函数，什么也不做
func void() {

}

func main() {
	// 声明 a 变量并打印
	var a int

	void()

	// 打印a变量的值和 dummy() 函数返回值
	fmt.Println(a, dummy(0))
}
