package welcome

import (
	"fmt"
	"time"
)


func SayHello() {
	fmt.Println("hello, world!")
}

func GetTime() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
}