package main

import (
	"fmt"
)

// interface{} 类型表示任意类型的值
type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// =============

// 为函数定义类型
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

// =============

func main() {
	// 声明接口
	var invoker Invoker
	// 实例化结构体
	s := new(Struct) // 或 &Struct
	// 将实例化的结构体赋值到接口
	invoker = s
	// 使用接口调用实例化的结构体方法 Struct.Call
	invoker.Call("hello")
}
