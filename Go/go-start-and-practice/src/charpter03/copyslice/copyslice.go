package main

import (
	"fmt"
)

func main() {
	const elementCount = 1000
	// 源切片
	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	// 引用切片
	refData := srcData

	// 拷贝切片
	copyData := make([]int, elementCount)
	copy(copyData, srcData)
	srcData[0] = 99
	// 查看引用切片是否发生变化
	fmt.Println(refData[0])
	// 查看拷贝切片是否发生变化
	fmt.Println(copyData[0], copyData[elementCount-1])
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
