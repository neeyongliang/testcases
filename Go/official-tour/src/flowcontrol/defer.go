package flowcontrol

import (
	"fmt"
)

func TestDefer() {
	defer fmt.Println("ahead but defer")
	fmt.Println("after")
}

