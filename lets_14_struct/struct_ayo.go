package main

import "fmt"

//GO语言中的匿名结构体 & 结构体嵌套

type dog struct {
	string
	int
}
type Person struct {
	name string
	age  int
	//add  address 可以这么写也可以
	address //这样写的好处是因为go语言有给语法糖  直接调用的了里面字段
}
type address struct {
	city    string
	zipcode int
}

func main() {
	d := dog{
		"小狗",
		2,
	}
	fmt.Println(d)
	//现在没有结构体没有字段名 我们就通过类型来调用
	fmt.Println(d.string)
	p := Person{
		name: "Ding",
		age:  21,
		address: address{
			city:    "ShangHai",
			zipcode: 200000,
		},
	}
	fmt.Println(p)
	fmt.Println(p.city)
}
