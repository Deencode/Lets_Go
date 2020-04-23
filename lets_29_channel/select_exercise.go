// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/23 - 8:33 下午

package main

import "fmt"

func main() {
	ints := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ints <- i: //第一个条件 通道缓冲区是1 有位置并且是空的可以发送值
		case x := <-ints: //第二个条件不满足,通道取值肯定是取不到的因为是空的
			fmt.Println(x)
		}
	}
}
