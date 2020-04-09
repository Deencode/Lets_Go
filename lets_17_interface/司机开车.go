package main

import "fmt"

type AbstractDrive interface { //司机抽象
	//driver()
	driveCar(a Auto)
}

type deriver struct {
	name string //司机姓名
}

type Auto struct { //汽车
	brand string
}

// func (d deriver) deriver() { //这种方法就是用来设置驾驶员
// 	fmt.Println("司机名字:", d.name)
// }
func (d deriver) driveCar(a Auto) { //设置开的什么 brand汽车
	fmt.Println("司机：", d.name, "开->", a.brand)
}
func main() {
	var absd AbstractDrive
	absd = deriver{
		name: "丁烁",
	}
	absd.driveCar(Auto{"法拉利_拉法"})
	absd = deriver{"岳宁宁"}
	absd.driveCar(Auto{"雷克萨斯LC500_美规版本"})
}

/*
	[Running] go run "d:\Go_workspace\src\Lets_Go\lets_17_interface\司机开车.go"
	司机： 丁烁 开-> 法拉利_拉法
	司机： 岳宁宁 开-> 雷克萨斯LC500_美规版本

	[Done] exited with code=0 in 5.305 seconds
*/
