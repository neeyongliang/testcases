package main

import "fmt"

type Bag struct {
	items []int
}

// 面向过程实现的方法
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

// 面向对象实现的方法
func (b *Bag) Insert(itemid int) {
	// 这里 (b *Bag) 就是一个接收器，每个方法只能有一个接收器
	// 官方建议是类型首字母的小写，其实随便些什么。。。
	b.items = append(b.items, itemid)
}

func main() {
	bag := new(Bag)
	Insert(bag, 1001)
	bag.Insert(1002)
	fmt.Println(bag.items)
}
