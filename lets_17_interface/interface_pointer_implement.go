package main

import "fmt"

type AbsDriver interface {
	driverCar(a Auto)
}

type Auto struct {
	brand string "json:Car_brand"
}

type driver struct {
	name string
}

//值接收者
// func (d driver) driverCar(a Auto) {
// 	fmt.Println(d.name+"开", a.brand)
// }

//指针接收者
func (d *driver) driverCar(a Auto) {
	fmt.Println(d.name+"开", a.brand)
}

/*
	值接收者实现的接口可以存储 结构体值类型的和结构体指针类型的都能存储
	指针接收者 只能存储结构体的指针类型
*/
func main() {
	var absd AbsDriver
	absd = &driver{"丁烁"} //这里有& 代表取地址 这个获取的就是结构体的指针类型实现 如果去掉“&”就变成了值类型结构体接收者就会报错
	absd.driverCar(Auto{"法拉利488"})
}
