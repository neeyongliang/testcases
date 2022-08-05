package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func test_len() {
	tip1 := "geni is a ninja"
	tip2 := "忍者"
	fmt.Println(len(tip1))
	fmt.Println(len(tip2))
}

func test_char_num() {
	tip1 := "geni is a ninja"
	tip2 := "忍者"
	fmt.Println(utf8.RuneCountInString(tip1))
	fmt.Println(utf8.RuneCountInString(tip2))
}

func test_string_connect() {
	hammer := "吃我一钯,"
	sickle := "俺老猪来也"
	// 声明字节缓冲
	var buf2 bytes.Buffer
	// 把字节写入缓冲
	buf2.WriteString(hammer)
	buf2.WriteString(sickle)
	fmt.Println(buf2.String())
}

func main() {
	test_len()
	test_char_num()
	test_string_connect()
}
