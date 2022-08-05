package methods

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs3() float64
}

type MyFloat2 float64

func (f MyFloat2) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex3) Abs3() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func TestInterfaces() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex3{3, 4}
	a = f
	fmt.Println(a.Abs3())
	a = &v
	fmt.Println(a.Abs3())
}

type I interface {
	M()
}

type T struct {
	S string
}

type O struct {
	P string
}
// 此方法表示类型 T 实现了结构 I， 但我们无需显示声明此事
func (t T) M() {
	fmt.Println(t.S)
}

func TestInterfacesSatisfiedImplicitly() {
	var i I = T{"Hello"}
	i.M()
}

type F float64
func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TestInterfacesValues() {
	var i I
	i = &T{"World"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func (o *O) M() {
	if o == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(o.P)
}

func TestInterfaceValueWithNil() {
	var i I
	var o *O

	i = o
	describe(i)
	i.M()

	i = &T{"zoo"}
	describe(i)
	i.M()
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TestEmptyInterface() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}

func TestTypeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) will raise a panic
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v * 2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TestTypeSwitches() {
	do(21)
	do("function")
	do(true)
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func TestStringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
