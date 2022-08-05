package methods

// go 没有类
import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestMethods() {
	v := Vertex3{
		X: 3,
		Y: 4,
	}
	fmt.Println(v.Abs())
}

func Abs2(v Vertex3) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func TestMethods2() {
	v := Vertex3{3, 4}
	fmt.Println(Abs2(v))
}

type MyFloat float64

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func TestMethods3() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs3())
}

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestMethodsPointers() {
	v := Vertex3{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

func Scale2(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestMethodsPointersExplained() {
	v := Vertex3{3, 4}
	Scale2(&v, 10)
	fmt.Println(Abs2(v))
}

func TestIndirection() {
	v := Vertex3{3, 4}
	v.Scale(2)
	Scale2(&v, 10)

	p := &Vertex3{4, 3}
	p.Scale(3)
	Scale2(p, 8)

	fmt.Println(v, p)
}

func TestIndirectionValues() {
	v := Vertex3{3, 4}
	p := &Vertex3{4, 3}
	// 以值为接收者的方法被调用时，接收者既能为值又能为指针
	fmt.Println(v.Abs())
	fmt.Println(p.Abs())

	// 接受一个值作为参数的函数必须接受一个指定类型的值
	fmt.Println(Abs2(v))
	fmt.Println(Abs2(*p))
}

func TestWithPointerReceivers() {
	v := &Vertex3{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %v, Abs: %v\n", v, v.Abs())
}
