package main

// map 为映射，表示键值关系
// 映射定义方式有两种：1，make 方式；2，声明之后加大括号方式

import (
	"fmt"
)

func test_map() {
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"]) // 输出 66

	v := scene["route2"]
	fmt.Println(v) // 输出 0
}

func test_map2() {
	m := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}
	fmt.Println(m["W"])
	v, ok := m["A"]
	fmt.Printf("Check key 'A' is exists: %v, value: %s\n", ok, v)
	delete(m, "A")
	v, ok = m["A"]
	fmt.Printf("Check again key 'A' is exists: %v, value: %s", ok, v)
}

func main() {
	test_map()
	test_map2()
}
