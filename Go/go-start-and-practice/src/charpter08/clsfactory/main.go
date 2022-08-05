package main

import (
	"charpter08/clsfactory/base"
	_ "charpter08/clsfactory/cls1" // 表示导入却不使用
	_ "charpter08/clsfactory/cls2" // 还是会调用各自的 init
)

func main() {
	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()
}