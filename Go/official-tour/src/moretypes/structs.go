package moretypes

import "fmt"

type Vertex struct {
	X int
	Y int
}

func TestStruct() {
	fmt.Println(Vertex{1, 2})
}

func TestStructFields() {
	v := Vertex{3, 4}
	fmt.Println(v.X)
}

func TestStructPointers() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // 等价于 (*p).X
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &Vertex{1, 2}
)

func TestStructLiterals() {
	fmt.Println(v1, p, v2, v3)
}