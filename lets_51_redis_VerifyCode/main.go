// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/8/20 - 6:09 下午 - UTC/GMT+08:00

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	phone  = "19999999999"
	format = "user:%s:code"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() {
	Rdb.SetNX(fmt.Sprintf(format, phone), RandomString(5), 60*time.Second)
}

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

//func randomCode(length int) (code []byte) {
//	seeds := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Println(seeds)
//	for i := 0; i < length; i++ {
//		fmt.Println(seeds[rand.Intn(9)])
//		code = append(code, seeds[rand.Intn(9)])
//	}
//	return
//}
