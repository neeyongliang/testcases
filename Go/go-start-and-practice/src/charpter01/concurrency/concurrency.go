package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数据生产者
func producer(header string, channel chan string) {
	// 无线循环，不停的产生数据
	for {
		// 讲随机数字字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("== %s: %v ==\n", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

// 数据消费者
func customer(channel chan string) {
	for {
		// 不停的获取数据
		message := <-channel
		fmt.Printf("Customer get message, details is :%s", message)
	}
}

func main() {
	// 创建一个字符串类型的通道
	channel := make(chan string)

	// 创建 producer() 函数的并发 goroutine
	go producer("cat", channel)
	go producer("dog", channel)

	// 消费数据
	customer(channel)
}
