package main

import "fmt"

//go语言中的bool布尔类型
var (
	b1 bool
	b2 = true
	b3 bool
)

func main() {
	fmt.Println(b1, b2, b3)
	fmt.Println(b1 == b3)
	fmt.Println(b2 == b3)
}

// [Running] go run "/Users/ding/Documents/GO_CODE_DEV/src/Lets_Go/lets_04_bool/main.go"
// false true false
// true
// false
// [Done] exited with code=0 in 0.232 seconds
