package main

import (
	"fmt"
)

type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

func main() {
	c := &C{}
	// 这里 c.a 会报错，不知道是哪个a
	c.A.a = 1
	fmt.Println(c.A.a)
}
