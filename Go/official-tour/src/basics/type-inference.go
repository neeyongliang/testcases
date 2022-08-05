package basics

import "fmt"

func TestTypeInfer() {
	v := 42 // fix here
	fmt.Printf("v is type %T\n", v)
}
