package main

import (
	"fmt"
)

// 用于测试值传递的结构体
type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}

// 代表各种结构体的字段
type InnerData struct {
	a int
}

// 值传递测试函数
func passByValue(inFunc Data) Data {
	// 输出参数成员
	fmt.Printf("inFunc value: %+v\n", inFunc)

	// 打印 inFunc 的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

func main() {
	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}

	fmt.Printf("in value: %+v\n", in)
	fmt.Printf("in ptr: %p\n", &in)
	out := passByValue(in)

	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)
}
