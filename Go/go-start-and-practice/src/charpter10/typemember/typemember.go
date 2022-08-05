package main

import (
	"fmt"
	"reflect"
)

func main() {
	type cat struct {
		Name string
		// 带有结构体tag的字段
		Type int `json:"type", id:"100"`
	}
	ins := cat{Name: "mimi", Type: 1}
	typeOfCat := reflect.TypeOf(ins)
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Printf("name:'%v', tag:'%v'\n", fieldType.Name, fieldType.Tag)
	}
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Printf(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
