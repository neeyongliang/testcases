package main

import (
	"fmt"
)

// 往关闭的通道里发送数据就会宕机
// 从关闭的通道里读取数据不再阻塞

func main() {
	ch := make(chan int)
	close(ch)
	fmt.Printf("ptr: %p, cap: %d len: %d\n", ch, cap(ch), len(ch))
	// 这一句会引发 panic，向已关闭的通道发送数据
	// ch <- 1

	ch2 := make(chan int, 2)
	ch2 <- 0
	ch2 <- 1
	close(ch2)

	for i := 0; i < cap(ch2); i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}
