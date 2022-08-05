package main

type Cat struct {
	Color string
	Name  string
}

// 模拟重载
func NewCatByColor(color string) {
	return &Cat{
		Color: color,
	}
}

func NewCatByName(name string) {
	return &Cat{
		Name: name,
	}
}

// 模拟父类
type BlackCat struct {
	Cat // 嵌入 Cat，类似于派生
}

// 构造父类
func newCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 构造子类
func newBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

func main() {

}
