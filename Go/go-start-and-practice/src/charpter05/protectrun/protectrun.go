package main

import (
	"fmt"
	"runtime"
)

type panicContext struct {
	function string
}

func protectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}

func main() {
	fmt.Println("run before")
	// 运行一段手动触发的错误
	protectRun(func() {
		fmt.Println("manual panic before")
		panic(&panicContext{"manual panic"})
		fmt.Println("manual panic after")
	})
	// 故意造成空指针访问的问题
	protectRun(func() {
		fmt.Println("assignment panic before")
		var a *int
		*a = 1
		fmt.Println("assignment panic after")
	})

	fmt.Println("run after")
}
