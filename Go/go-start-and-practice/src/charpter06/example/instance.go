package main

type Point struct {
	X int
	Y int
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

type Command struct {
	Name    string
	Var     *int
	Comment string
}

var version int = 1

// newCommond 是方法 3 的一个封装
func newCommond(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

func main() {
	// 1. 基本实例化形式
	var p Point
	p.X = 4
	p.Y = 3

	// 2. 使用 new 实例化
	// 这里 Go 语言使用了 语法糖（Syntactic sugar）
	// inst.Name <=> (*inst).Name
	tank := new(Player)
	tank.Name = "Cannon"
	tank.HealthPoint = 300
	tank.MagicPoint = 0

	// 3. 使用取地址符 & 实例化
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"

	// 使用 newCommand 封装实现方法 3
	cmd = newCommond("version", &version, "show version")
}
