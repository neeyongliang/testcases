package main

import "os"
import "fmt"

// usage:
// $ go build $FILE
// $ ./build_obj a b c d
// return:
// [./59command-line-arguments]
// [a b c d]
// c

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
