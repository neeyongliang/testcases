package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Server

func server(address string, exitChan chan int) {
	// 监听给定的地址
	l, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println(err.Error())
		exitChan <- -1
	}

	fmt.Println("listen: " + address)

	defer l.Close()

	for {
		// 新链接没有到来时，Accept 是阻塞的
		// conn 的类型为 *tcp.Conn
		conn, err := l.Accept()

		// 连接失败
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go handleSession(conn, exitChan)
	}
}

// Session
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started...")
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			str = strings.TrimSpace(str)
			// 处理 telnet 命令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			// 发送回音：收到什么发什么
			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println("Session close...")
			conn.Close()
			break
		}
	}
}

// TelnetCommand
func processTelnetCommand(str string, exitChan chan int) bool {
	// 遇到 @close 表示终止会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed...")
		// 告诉外部要断开连接
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown...")
		exitChan <- 0
		// 告诉外部要断开连接
		return false
	}
	fmt.Println(str)
	return true
}

func main() {
	exitChan := make(chan int)
	go server("127.0.0.1:7001", exitChan)

	// 通道阻塞，等待范湖结果
	code := <-exitChan

	// 记录返回值并推出
	os.Exit(code)
}
