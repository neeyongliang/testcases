package main

import (
	"fmt"
)

func testSwitch() {
	// Go 语言每个 case 自带了 break，不会向下继续执行
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

}

func testSwitch2() {
	// 一个分支，多个值
	var a = "mum"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	}
}

func testSwitch3() {
	// 无条件 switch
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}
}

func testSwitch4() {
	// 兼容 C 语言的设计：fallthrough，新编写的代码不要使用
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello, ")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}

func main() {
	testSwitch()
	testSwitch2()
	testSwitch3()
	testSwitch4()
}
