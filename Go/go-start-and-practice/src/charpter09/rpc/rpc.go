package main

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string, req string) (string, error) {
	ch <- req
	// 使用 select 做多路复用，与 switch 相比，select 无法接任何语句
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	for {
		// 接收客户端数据
		data := <-ch
		// 打印数据
		fmt.Println("server received:", data)

		// 通过睡眠函数可以出发超时数据
		time.Sleep(time.Second * 2) // 1 秒钟超时，现在暂停 2 秒

		// 向客户端反馈已收到
		ch <- "roger"
	}
}

func main() {
	ch := make(chan string)
	go RPCServer(ch)
	recv, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received", recv)
	}
}
