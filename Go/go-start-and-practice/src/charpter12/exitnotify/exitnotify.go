package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// 使用 goroutine
func socketRecvByChannel(conn net.Conn, exitChan chan string) {
	buff := make([]byte, 1024)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			break
		}
	}
	exitChan <- "recv exit"
}

func goRoutine(conn net.Conn) {
	exit := make(chan string)
	go socketRecvByChannel(conn, exit)
	time.Sleep(time.Second)
	conn.Close()
	fmt.Println(<-exit)
}

// 使用 WaitGroup
func socketRecvByWaitGroup(conn net.Conn, wg *sync.WaitGroup) {
	buff := make([]byte, 1024)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			break
		}
	}
	wg.Done()
}

func waitGroup(conn net.Conn) {
	var wg sync.WaitGroup
	wg.Add(1)
	go socketRecvByWaitGroup(conn, &wg)
	time.Sleep(time.Second)
	conn.Close()
	wg.Wait()
	fmt.Println("recv done")
}

func main() {
	conn, err := net.Dial("tcp", "www.163.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用 goroutine
	goRoutine(conn)

	// 使用 WaitGroup
	waitGroup(conn)
}
