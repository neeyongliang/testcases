package main

import (
	"fmt"
)

func printer(c chan int) {
	for {
		data := <-c
		if data == 0 {
			break
		}

		fmt.Println(data)
	}
	// 通知 main 函数已经结束循环（我搞定了）
	c <- 0
}

func main() {
	c := make(chan int)
	go printer(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}

	// 通知并发的 printer 已经结束循环（没数据啦）
	c <- 0

	// 等待 printer 结束（搞定了喊我）
	<-c
}
