package main

import "fmt"

//go语言中的defer关键字的使用
func main() {
	/*
		go语言中的defer关键字表示的代码会放到程序其他代码执行完毕之后在执行
		并且defer是执行顺序是倒序
		以下代码运行结果:
			start -> end ->3-> 2-> 1
	*/
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")

	fmt.Println(f1())       // 5
	fmt.Println(f2())       // 6
	fmt.Println(f3())       // 5
	fmt.Println(f4())       // 5
	fmt.Println("f5", f5()) // 6
}

//defer经典面试题
// 过程: x = 5 然后 x赋值给ret  然后执行 defer defer里面方法是修改之前在函数体里面的x没有修改返回的
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//这个是在func返回值就设置了int类型的x作为返回值
//5赋值给x 然后 执行defer 里面的x++ 此时这个x是function里面的副本
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

//定义变量x =5 然后把 x 赋值给 y 然后执行defer 但是defer里面的x的值是x的  返回值是y的 互不干扰
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//5赋值给x 然后执行defer defer的func把x作为参数传入进去了
//然后修改只是内存中的副本 返回值早已经在第一步已经赋值操了 互不干扰
func f4() (x int) {
	defer func(x int) { // 有参数只是给副本
		x++
	}(x)
	return 5
}

/*
	1. 函数的的返回值已经定义一个返回值变量x
	2. 然后把5赋值给x x=5
	3. 然后执行defer
	4. defer匿名函数把x指针传入进去了
	5. 然后通过指针地址获取到了x存储的值
	6. 然后把x的值加1 x = 5 + 1 ——> 6
	7. 次数这个x的指针和返回值ret存储那个x的指针地址一样 所以就达到了修改作用
	8. 返回了
*/
func f5() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5
}
