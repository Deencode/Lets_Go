// GO语言中的函数编程Example
// http://github.com/higker/lets_go
package main

import "fmt"

//go语言中的闭包操作
func main() {
	f1(f3(f2, 5, 5))
}

//1.限制函数类型的函数
func f1(f func()) {
	f()
}

//2.现在的需求就要把我们的f2函数传递到f1函数里面进行调用
func f2(a, b int) int {
	fmt.Println("f2() am f1() exec succeed")
	fmt.Println("a + b = ? 在f3中的匿名函数中调用:")
	return a + b
}

//3.通过闭包解决
func f3(fn func(int, int) int, a, b int) (Rf func()) { //返回值名字必须是“Rf” 不想用就不需要写直接写返回值类型
	//这个就是取一个中间变量存储 并且 在函数里面声明一个匿名无返回值的函数
	// := 这里如果函数返回值已经在函数名上写着了 就不能使用 := 因为已经在创建函数时已经声明了
	Rf = func() {
		sum := fn(a, b) //接受返回值 然后打印
		fmt.Println("sum=", sum)
	}
	return Rf
}
