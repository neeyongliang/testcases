package main

import (
	"fmt"
)

type Alipay struct {
}

func (a *Alipay) CanUseFaceID() {
}

type Cash struct {
}

func (c *Cash) Stolen() {

}

// 具备刷脸特性
type CantainCanUseFaceID interface {
	CanUseFaceID()
}

// 具备被偷特性
type ContainStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceID:
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}

func main() {
	print(new(Alipay))
	print(new(Cash))
}
