package main

import (
	"fmt"
	"time"
)

// 一般格式，var 通道名 chan 通道类型
// 需要配合 main

func main() {
	ch := make(chan int)
	go func() {
		for i := 3; i >= 0; i-- {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for data := range ch {
		fmt.Println(data)
		if data == 0 {
			break
		}
	}
}
