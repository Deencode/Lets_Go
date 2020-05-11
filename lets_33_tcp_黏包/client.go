// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/11 - 7:58 下午

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:9598")
	if err != nil {
		fmt.Println("connection tcp server fail.", err)
	}
	for i := 0; i < 10; i++ {
		dial.Write([]byte("Hello" + time.Now().Format("2006-01-02 15:04:05.0000")))
	}
}
