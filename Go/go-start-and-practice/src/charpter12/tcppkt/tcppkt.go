package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
)

// 二进制封包格式
type Packet struct {
	Size uint16
	Body []byte
}

// 将数据写入 dataWriter
func WritePacket(dataWriter io.Writer, data []byte) error {
	// 准备一个缓冲字节
	var buffer bytes.Buffer
	// 将 Size 写入缓冲，将 uint16 格式传入的切片大小，以小端方式写入缓冲中
	if err := binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil {
		return err
	}

	if _, err := buffer.Write(data); err != nil {
		return err
	}

	// 获得完整的数据
	out := buffer.Bytes()
	if _, err := dataWriter.Write(out); err != nil {
		return err
	}

	return nil
}

// 连接器，传入连接地址和发送封包的次数
func Connector(address string, sendTimes int) {
	// 尝试用 Socket 连接地址
	conn, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < sendTimes; i++ {
		str := strconv.Itoa(i)
		if err := WritePacket(conn, []byte(str)); err != nil {
			fmt.Println(err)
			break
		}
	}
}

// 接受器
type Acceptor struct {
	// 保存侦听器
	l net.Listener
	// 侦听器的停止同步
	wg            sync.WaitGroup
	OnSessionData func(net.Conn, []byte) bool
}

// 异步开始侦听
func (a *Acceptor) Start(address string) {
	go a.listen(address)
}

func (a *Acceptor) listen(address string) {
	a.wg.Add(1)
	defer a.wg.Done()
	var err error
	a.l, err = net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		// 新连接到来时，Accept 是阻塞的
		conn, err := a.l.Accept()
		if err != nil {
			break
		}
		go handleSession(conn, a.OnSessionData)
	}
}

func (a *Acceptor) Stop() {
	a.l.Close()
}

func (a *Acceptor) Wait() {
	a.wg.Wait()
}

func NewAcceptor() *Acceptor {
	return &Acceptor{}
}

// 从 dataReader 中读取封包
func readPacket(dataReader io.Reader) (pkg Packet, err error) {
	var sizeBuffer = make([]byte, 2)
	_, err = io.ReadFull(dataReader, sizeBuffer)
	if err != nil {
		return
	}

	// 使用 bytes.Reader 读取 sizeBuffer 中的数据
	sizeReader := bytes.NewReader(sizeBuffer)
	// 读取小端 unit16 作为 size
	err = binary.Read(sizeReader, binary.LittleEndian, &pkg.Size)
	if err != nil {
		return
	}

	pkg.Body = make([]byte, pkg.Size)
	_, err = io.ReadFull(dataReader, pkg.Body)
	return
}

// 连接会话逻辑
func handleSession(conn net.Conn, callback func(net.Conn, []byte) bool) {
	dataReader := bufio.NewReader(conn)
	for {
		pkg, err := readPacket(dataReader)
		if err != nil || !callback(conn, pkg.Body) {
			conn.Close()
			break
		}
	}
}

func main() {
	const TestCount = 100000
	const address = "127.0.0.1:8010S"
	var recvCounter int
	acceptor := NewAcceptor()
	acceptor.Start(address)
	acceptor.OnSessionData = func(conn net.Conn, data []byte) bool {
		str := string(data)
		n, err := strconv.Atoi(str)
		if err != nil || recvCounter != n {
			panic("failed")
		}
		recvCounter++
		if recvCounter >= TestCount {
			acceptor.Stop()
			return false
		}
		return true
	}
	Connector(address, TestCount)
	acceptor.Wait()
}
