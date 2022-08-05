package gotest

import (
	"fmt"
	"testing"
)

// 测试框架提供的日志方法：
// Log, Logf
// Error, Errorf
// Fatal, Fatalf

func TestFail(t *testing.T) {
	fmt.Println("before fail")
	t.Fail()
	fmt.Println("after fail")
}
