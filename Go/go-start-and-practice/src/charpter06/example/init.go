package main

// 爸爸的爸爸叫爷爷
type People struct {
	name  string
	child *People
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

func main() {
	// 1. 键值对的初始化形式
	// 结构体成员中，只能包含结构体类型的指针
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	// 2. 无键值初始化
	// 必须初始化所有字段；不可与键值类型混用；顺序必须正确
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}

}
