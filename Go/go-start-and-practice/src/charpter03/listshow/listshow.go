package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("canon")
	l.PushFront(67)
	element := l.PushBack("fist")
	l.InsertBefore("hight", element)
	l.InsertBefore("none", element)
	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
