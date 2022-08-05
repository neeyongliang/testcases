package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 两个并发程序间的同步：
// 1. 通道
// 2. 互斥锁
// 3. 等待组

var (
	seq   int64
	count int
	// 除了互斥锁意外，还有读写互斥锁 sync.RWMutex, RLock, RUnlock
	countGuard sync.Mutex
)

// 使用 go run -race xxx.go 来检测竞态问题

func GenID() int64 {
	// 这两句会报竞态问题的错
	// atomic.AddInt64(&seq, 1)
	// return seq
	// 原子操作比 sync.Mutex 更高效
	return atomic.AddInt64(&seq, 1)
}

func GetCount() int {
	countGuard.Lock()
	defer countGuard.Unlock()
	return count
}

func SetCount(c int) int {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
	return count
}

func main() {
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID())

	SetCount(1)
	fmt.Println(GetCount())
}
