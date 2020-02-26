package main

import "fmt"

func main() {
	n := "100"
	// var ptr *int
	i := 200
	ptr := &i
	fmt.Println("n 的pointer", ptr)
	fmt.Println("i=", i)
	n = fmt.Sprintf("%v", *ptr)
	fmt.Println("n=", n)
	//fmt.Println(&n == &i)
}
