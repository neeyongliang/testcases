package main

import (
	"fmt"
)

type ChipType int

// go 语言目前没有枚举，使用整型配合 iota 模拟枚举
const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	fmt.Printf("%s %d", CPU, CPU)
}
