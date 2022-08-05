package moretypes

import (
    "fmt"
    "strings"
)


func TestSlicesPointers() {
	names := [4]string {
		"Jhon",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func TestSliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func TestSliceBounds() {
	s := []int{2, 3, 4, 5, 6, 7}
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func TestPrintSlice() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// 截取切片使其长度为 0
	s = s[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// 拓展长度
	s = s[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// 舍弃前两个值
	s = s[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func TestNilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func TestMakingSlices() {
	a := make([]int, 5) // len(a)=0, cap(a)=5
	fmt.Println("a len=%d cap=%d %v\n", len(a), cap(a), a)
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println("b len=%d cap=%d %v\n", len(b), cap(b), b)
	c := b[:2]
	fmt.Println("c len=%d cap=%d %v\n", len(c), cap(c), c)
	d := c[2:5]
	fmt.Println("d len=%d cap=%d %v\n", len(d), cap(d), d)
}

func TestSlicesOfSlice() {
    board := [][]string{
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
    }
     board[0][0] = "X"
     board[2][2] = "O"
     board[1][2] = "X"
     board[1][0] = "O"
     board[0][2] = "X"

     for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
     }
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func TestSliceAppend() {
	var s []int
	printSlice(s)
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	s = append(s, 2, 3, 4)
	printSlice(s)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func TestRange() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func TestRangeContinue() {
	pow := make([]int, 10)
	for i := range pow { // 等价于 for i, _ := range pow
		pow[i] = 1 << uint(1) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}