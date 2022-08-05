package main

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	Zero Enum = 0
)

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	// 类型，类属（比如指针）
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	type cat struct {
	}

	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	typeOfB := reflect.TypeOf(Zero)
	fmt.Println(typeOfB.Name(), typeOfB.Kind())

	ins := &cat{}
	typeOfD := reflect.TypeOf(ins)
	fmt.Printf("name:'%v', kind:'%v'\n", typeOfD.Name(), typeOfD.Kind())
	typeOfE := typeOfD.Elem()

	fmt.Printf("element name:'%v', element kind:'%v'\n", typeOfE.Name(), typeOfE.Kind())
}
