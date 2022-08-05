package main

import (
	"bytes"
	"fmt"
)

// ...string <=> []stringS
func joinString(slist ...string) string {
	var b bytes.Buffer
	// 使用 WriteString 为了更高效的合并字符串
	for _, s := range slist {
		b.WriteString(s)
	}

	return b.String()
}

func main() {
	fmt.Println(joinString("pig ", "and ", "cat"))
}
