package main

import (
	"fmt"
)

type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

type Walkable struct{}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

type Human struct {
	Walkable
}

type Bird struct {
	Walkable
	Flying
}

func main() {
	b := new(Bird)
	fmt.Println("Bird: ")
	b.Fly()
	b.Walk()

	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()
}
