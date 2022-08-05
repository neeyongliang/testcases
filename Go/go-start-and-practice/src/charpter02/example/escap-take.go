package main

// go run -gcflags "-m -l" xxx.go

import "fmt"

// Data 声明空的结构体测试变量逃逸的情况
type Data struct {
}

func dummy2() *Data {
	// 实例化 c 为 Data 类型
	var c Data
	// 返回 c 的地址
	return &c
}

func main() {
	fmt.Println(dummy2())
}

// Note:
// 编译器决定变量分配在堆还是栈上，取决于：
// 1. 变量是否被取地址
// 2. 变量是否发生逃逸
