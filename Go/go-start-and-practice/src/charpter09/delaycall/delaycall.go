package main

import (
	"fmt"
	"time"
)

func main() {
	// 退出用的通道
	exit := make(chan int)

	fmt.Println("start")
	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")
		// 通知 main 的 goroutine 已经结束
		exit <- 0
	})

	fmt.Println("wait exit")
	// 等待结束，这里会被阻塞住
	<-exit
}
