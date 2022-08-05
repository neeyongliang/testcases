package main

import (
	"fmt"
)

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func main() {
	e := New("custom error")
	fmt.Println(e.Error())
}
