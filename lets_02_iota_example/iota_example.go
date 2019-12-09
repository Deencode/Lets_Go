package main

import "fmt"

//iota 是go语言中的常量计数器
//iota会在const关键字出现的时候 iota的值虎重置为0
const (
	n1 = iota //0

	n2 // 1

	n3 // 2
)

//测试 iota 的作用
const (
	A1 = 1
	A2 = iota
	A3
)

//一道题出给自己做的
const (
	v1 = 2
	v2 = 3
	v3 = iota
	v4 = iota
	v5
	v6 = iota
	v7
	//v8 = 8 + 1
	v8
)

const (
	c1 = 1       //value = 1 iota = 0
	c2           //value = 1 iota = 1
	c3 = iota    // value = 2  iota = 2
	c4           // vlaue = 3  iota = 3
	c5 = "Hello" //vlaue = Hello   iota = 4
	c6           //value = Hello   iota = 5
	c7           //value = Hello   iota = 6
	c8 = iota    //value = 7  iota = 7
	//1 1 2 3 Hello Hello Hello  7
)

func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(A1, A2, A3)
	//个人预测结果是 2 3 2 3 4 5 6 8
	// const 常量在运行时内存是改变不了的 所以原来的值是不变的
	// 但是iota关键字会统计常量的个数   个人感觉 iota应该在内存不断给没有的赋值的const的常量在赋值
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8) //2 3 2 3 4 5 6 8
	//1 1 2 3 Hello Hello Hello  7
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8)
	/*
	 * 总结一下:
	 * iota关键字在const出现时会重置为零
	 * iota会统计const块里面的常量的个数
	 * 每在const块里定义一个常量 iota的值就会累加1
	 */
	//和我预测的结果一样  掌声👏给自己鼓励一下
}
