package gotest

import "testing"

func TestA(t *testing.T) {
	t.Log("A")
}

func TestAK(t *testing.T) {
	t.Log("AK")
}

func TestB(t *testing.T) {
	t.Log("B")
}

func TestC(t *testing.T) {
	t.Log("C")
}

// go test -v -run TestA select_test.go
// -run 跟随的测试用例名，支持曾泽表达式，所以会执行 TestA, TestAK
// 可以使用 -run TestA$ 执行单个
