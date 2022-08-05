package main

import (
	"fmt"
	"reflect"
)

type Brand struct {
}

func (t Brand) Show() {
	fmt.Println("Brand show...")
}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {
	var a Vehicle
	a.FakeBrand.Show()
	ta := reflect.TypeOf(a)
	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		fmt.Printf("Field Name: %v, FieType: %v\n", f.Name, f.Type.Name())
	}
}
