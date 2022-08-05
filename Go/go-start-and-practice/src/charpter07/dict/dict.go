package main

import "fmt"

type Dictionary struct {
	// 键值都是任意数据类型
	data map[interface{}]interface{}
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}

	// callback 返回 true 时继续遍历
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

// 清空所有数据，就是 make 而已，没有其他操作
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}

func main() {
	dict := NewDictionary()
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)

	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)

	// 这两处 return true 时为了循环能够一直执行下去
	dict.Visit(func(key, value interface{}) bool {
		if value.(int) > 40 {
			fmt.Println(key, "is expensive")
			return true
		}
		fmt.Println(key, "is cheap")
		return true
	})
}
