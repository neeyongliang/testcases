package flowcontrol

import "fmt"

func TestDeferMulti() {
	fmt.Println("TestDeferMulti start")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("TestDeferMulti done")
}
