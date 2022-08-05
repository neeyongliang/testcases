package main

import (
	"fmt"
)

type service interface {
	// 方法首字母大写才能被外部调用
	Start()     // 开启服务
	Log(string) // 日志输出
}

type Logger struct {
}

// Logger 实现了 Log 方法
func (g *Logger) Log(l string) {
	fmt.Println(l)
}

type gameService struct {
	Logger // 内嵌 Logger，“集成”它的方法
}

// 实现 Serivce 的 start() 方法
func (g *gameService) Start() {

}

func main() {
	var s service = new(gameService)
	s.Start()
	s.Log("ooooooo")
}
