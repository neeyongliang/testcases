package concurrency

import (
	"fmt"
	"golang-tour/mytree"
)

// Walk 补进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *mytree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *mytree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		fmt.Println(x, y)
		if x != y {
			return false
		}
	}
	return true
}

func ExerciseEqBinaryTree() {
	fmt.Println(Same(mytree.New(1), mytree.New(1)))
}
