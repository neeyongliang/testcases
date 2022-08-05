package main

import (
	"fmt"
)

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("Write Data:", data)
	return nil
}

func main() {
	f := new(file)
	var writer DataWriter
	// 此处变量类型不一致，但 writer 是一个接口，
	// 且 f 实现了 DataWriter 所有的方法
	writer = f
	writer.WriteData("data")
}
