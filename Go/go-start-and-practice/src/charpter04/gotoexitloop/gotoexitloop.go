package main

import "fmt"

func testNoGoTo() {
	var breakAgain bool
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				breakAgain = true
				break
			}
		}
		if breakAgain {
			break
		}
	}
	fmt.Println("no go to ... done")
}

func testGoTo() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	return

breakHere:
	fmt.Println("go to ... done")
}

func testErrorHandler() {
	// 使用 GoTo 进行错误统一处理
	err := firstCheckError()
	if err != nil {
		goto onExit
	}

	err = secondCheckError()
	if err != nil {
		goto onExit
	}

onExit:
	fmt.Println(err)
	exitProcess()
}

func main() {
	testNoGoTo()
	testGoTo()
}
