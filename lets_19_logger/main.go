package main

import (
	"Lets_Go/lets_19_logger/log"
)

const logPath string = "D:/Go_workspace/src/Lets_Go/lets_19_logger/log.log"

//记录日志测试
func main() {
	test()
	test1()
	test2()
}

func test1() {
	logs := log.NewFlog(logPath)
	//fmt.Printf("%T",logs)
	logs.Info("测试info")
	logs.DeBug("测试Debug")
	logs.Error("测试错误121212121")
}
func test2() {
	logs := log.NewClog()
	//fmt.Printf("%T",logs)
	logs.Info("测试info")
	logs.DeBug("测试Debug")
	logs.Error("测试错误错误错误错误错误")
}
func test() {
	f := &log.Flog{logPath}
	log.Logger(f).DeBug("测试")
}
