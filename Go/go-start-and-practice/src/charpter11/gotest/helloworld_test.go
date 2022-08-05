package gotest

import "testing"

// 1. 必须名必须以 _test.go 结尾
// 2. 函数名必须满足定义 func TestXXX(t *testing.T) 的定义
// 3. 不需要 main 函数，加上 -v 参数可以看到细节

func TestHelloWrold(t *testing.T) {
	t.Log("hello, world")
}
