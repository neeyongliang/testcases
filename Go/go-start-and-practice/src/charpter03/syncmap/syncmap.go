package main

// Go 语言中，map 在并发状态下读线程安全，同时写线程不安全
// Go 语言在1.9版本提供了 sync.Map
// Store: 存储，Load：读取，Delete: 删除
// 使用 Range 配合回调进行遍历

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map

	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	v, ok := scene.Load("london")
	fmt.Println(v, ok)
	scene.Delete("london")
	fmt.Println("-------------")
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("interface:", k, v)
		return true
	})
}
