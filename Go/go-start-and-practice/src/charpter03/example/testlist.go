package main

import "container/list"

// list 初始化的两种方法：
// 变量名 := list.New()
// var 变量名 list.List

func main() {
	l := list.New()
	l.PushBack("first")
	l.PushFront(67)
}
