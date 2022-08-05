package main

import (
	"fmt"
)

// 使用 MyImplement 实现 fmt.Stringer 方法

type MyImplement struct{}

func (m *MyImplement) String() string {
	return "hi"
}

func GetString() fmt.Stringer {
	var s *MyImplement = nil
	// 此处返回的 s 并不等于 nil，因为其中包含 type 信息
	return s
}

func main() {
	if GetString() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
		fmt.Printf("%v\n", GetString())
	}
}
