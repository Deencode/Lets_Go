package main

import "fmt"

//decimal go语言的进制转换
//二进制 八进制  十进制  十六进制

var (
	//十进制  PS:在go语言中无法直接给变量赋值为二进制的值
	//因为go语言有类型推断 会把你书写的二进制的数字 推断为int类型
	i1 = 101
)

func main() {
	fmt.Printf("i1的十进制是=%d\n", i1)
	fmt.Printf("i1的二进制进制是=%b\n", i1)
	fmt.Printf("i1的八进制进制是=%o\n", i1)
	fmt.Printf("i1的十六进制进制是=%x\n", i1)

	//查看变量类型
	fmt.Printf("i1的数据类型是%T\n", i1)
	i1 := int8(110)
	fmt.Printf("i1的数据类型是%T\n", i1)
}

// [Running] go run "/Users/ding/Documents/GO_CODE_DEV/src/Lets_Go/lets_03_decimal/main.go"
// i1的十进制是=101
// i1的二进制进制是=1100101
// i1的八进制进制是=145
// i1的十六进制进制是=65
// i1的数据类型是int
// i1的数据类型是int8
