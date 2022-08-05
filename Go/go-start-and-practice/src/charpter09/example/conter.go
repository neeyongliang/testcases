package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

func main() {
	// 使用 go 关键字创建的 goroutine 会忽略返回值
	// go running()

	// 也可以使用匿名函数
	go func() {
		var times int
		for {
			times++
			fmt.Println("dida", times)
			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
}
