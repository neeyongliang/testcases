package main

import (
	"flag"
	"fmt"
)

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	// 直接调用
	func(data int) {
		fmt.Println("hello,", data)
	}(100)

	// 赋值给函数
	f := func(data int) {
		fmt.Println("yoooho,", data)
	}

	f(9527)

	// 作为回调函数
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println("visit=>", v)
	})

	// 实现操作封装
	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
