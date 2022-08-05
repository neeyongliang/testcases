package main

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string

	// 嵌入字段
	float32
	bool
	next *dummy
}

func main() {
	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})

	fmt.Println("Num Field:", d.NumField())

	// 获取索引为 2 的字段
	floatField := d.Field(2)
	fmt.Println("Field:", floatField.Type)

	fmt.Println("Field", floatField.Type())

	fmt.Println("Field By Name(\"b\").Type:", d.FieldByName("b").Type())
	// []int{4,0} 表示 dummy 第四个元素（dummy类型）的第一个值，int型
	fmt.Println("Field By Index([]int{4,0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
}
