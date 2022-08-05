package main

import "fmt"

type Property struct {
	value int
}

func (self *Property) SetValue(v int) {
	self.value = v
}

func (self *Property) GetValue() int {
	return self.value
}

type Point struct {
	X int
	Y int
}

func (self Point) Add(other Point) Point {
	return Point{self.X + other.X, self.Y + other.Y}
}

func main() {
	// 指针型接收器
	p := new(Property)
	p.SetValue(100)
	fmt.Println(p.GetValue())

	// 非指针型接收器, 类似于只读方法，内部不会对成员变量进行修改
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	result := p1.Add(p2)
	fmt.Println(result)

	// 由于赋值的速度问题，小对象适合使用非指针接收器，大对象适合指针接收器
}
