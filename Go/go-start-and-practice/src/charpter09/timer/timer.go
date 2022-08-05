package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个 打点器，每 500 毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 5)
	// 创建一个计时器，两秒后触发
	stopper := time.NewTimer(time.Second * 2)
	// NewTicker, NewTimer 结构体内部，有个名为 C 的 <-chan
	// 类型的通道

	var i int
	for {
		select {
		case <-stopper.C:
			// 计时器时间到了
			fmt.Println("stop")
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick", i)
		}
	}

StopHere:
	fmt.Println("done")
}
