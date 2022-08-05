package main

import (
	"fmt"
	"strings"
)

func StringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

// 步骤：
// 1. 准备要处理的字符串
// 2. 准备字符串处理链
// 3. 处理字符串处理链
// 4. 打印输出后的字符串列表

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go fomater",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProcess(list, chain)
	for _, str := range list {
		fmt.Println(str)
	}
}
