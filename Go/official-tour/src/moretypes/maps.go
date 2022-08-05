package moretypes

import (
	"fmt"
)

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func TestMaps() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		Lat:  40.48433,
		Long: -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

var n = map[string]Vertex2 {
	"Bell Labs": Vertex2{
		40.48433,
		-74.39967,
	},
	"Google": Vertex2{
		37.42202,
		-122.08408,
	},
}

func TestMapLiterals() {
	fmt.Println(n)
}

// 若顶级类型只是一个类型名，你可以在文法的元素中省略它
var o = map[string]Vertex2 {
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func TestMapsLiteralsContinue() {
	fmt.Println(o)
}

func TestMutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:",v, "Present?", ok)
}