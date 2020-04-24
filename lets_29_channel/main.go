// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/22 - 3:15 下午

package main

import (
	"fmt"
	"sync"
)



var wg sync.WaitGroup
//go语言的channel 通道
//go语言使用goroutine执行多个task任务
//操作一个map或者其他变量或者内存时会发生数据竞争
//使用go里面使用channel来解决并发并行来解决
func main() {
	//channel类型 相当于一个队列 first in -> first out 先进先出原则
	//定义一个channel类型 channel是一个引用类型 需要开辟空间
	var ch chan int
	fmt.Println(ch == nil) //true
	//初始化一个无缓冲区的通道
	//把ch引用指向一个有内存地址的
	ch = make(chan int)
	wg.Add(1)
	go func (){
		defer wg.Done()
		x := <- ch
		fmt.Println(x)
	}()
	ch <- 2048 //会卡死main的主goroutine从而不能让程序进行执行
	fmt.Println(ch) //无缓冲区的channel不能放入值
	fmt.Println(ch == nil) //false
	//初始化一个带缓冲区的通道
	bufChan := make(chan int,8)
	fmt.Println(cap(bufChan)) //8 CAP可以换取通道缓冲区大小
	//发送值 使用这个"<-符号" 获取也一样
	bufChan <- 1024
	num := <-bufChan
	fmt.Println(cap(bufChan))
	fmt.Println(bufChan)
	fmt.Println(num)
	wg.Wait()
}


