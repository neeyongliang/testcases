package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i:= 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func TestGoroutines() {
	go say("world")
	say("???")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TestChannels() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收
	fmt.Println(x, y, x + y)
}

func TestBufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch<-3 will raise deadlock if buffer is overflow
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func TestRangeAndClose() {
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		print("+++++++++++++++")
		// select 会阻塞到可以执行位置
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		fmt.Println("---------------")
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci3(c, quit)
}

func TestDefaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// SafeCounter 的并发是安全的
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定的 key 的计数器的当前值
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后，同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func TestMutexCounter() {
	c := SafeCounter{v: make(map[string]int)}
	for i :=0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}