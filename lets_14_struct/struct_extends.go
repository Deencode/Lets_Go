package main

import "fmt"

//go语言中的继承
type animal struct {
	name string
	age  int
}
type tom struct {
	animal //这个就代表集成 有animal里面的所有方法
}

func (t tom) say() {
	fmt.Println(t.name, "喵喵~")
}

func main() {
	tom := tom{
		animal{ //这就是go里面的继承
			name: "小黄",
			age:  3,
		},
	}
	tom.say()
}
