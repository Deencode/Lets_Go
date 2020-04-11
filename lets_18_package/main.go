package main

//go语言导包
import (
	m "Lets_Go/lets_18_package/math" //起别名
	"fmt"
)

func main() {
	multip := m.SimpleCompute("+")
	fmt.Println(multip(2, 8))
}
