package main

import (
	"fmt"
	"sort"
)

// 将[]string 定义为 MyStringList
type MyStringList []string

// 实现 sort.Interface 接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Swap(i, j int) {
	// 这里是切片！
	m[i], m[j] = m[j], m[i]
}

func main() {
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1, First Blood",
	}

	// 其实 sort 包中自带了 sort.intSlice, sort.StringSlice
	sort.Sort(names)
	for _, name := range names {
		fmt.Println(name)
	}
}
