package cls2

import (
	"charpter08/clsfactory/base"
	"fmt"
)

type Class2 struct {
}

func (c *Class2) Do() {
	fmt.Println("Class2")
}

func init() {
	base.Register("Class2", func() base.Class {
		return new(Class2)
	})
}
