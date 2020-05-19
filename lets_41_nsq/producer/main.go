// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/19 - 9:17 下午

package main

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

var (
	// 公共的一个producer
	pc   *nsq.Producer
	adrr = "127.0.0.1:4150" //这个是nsqd 的tcp地址和port
)

func init() {
	cofing := nsq.NewConfig()                      // new配置
	producer, err := nsq.NewProducer(adrr, cofing) // new 一个消费者
	if err != nil {
		fmt.Println(" create producer failed.ERR =", err)
		return
	}
	pc = producer
}

type student struct {
	Types           interface{}
	TypesStructName string
	Id              int
	Name            string
	Age             int
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("producer star....")
	for {
		//str, _ := reader.ReadString('\n')
		err := pc.Publish("topic_01", []byte(randData()))
		if err != nil {
			fmt.Println("push topic_01 message failed.ERR=", err)
		}
		fmt.Println("send successful msg.")
		//time.Sleep(1 * time.Second)
	}
}
func randData() string {
	rand.Seed(time.Now().UnixNano())
	kind := reflect.TypeOf(student{}).Kind()
	stu := &student{Types: kind, TypesStructName: "student", Id: rand.Intn(100), Name: "学生" + strconv.Itoa(rand.Intn(100)), Age: rand.Intn(100)}
	marshal, err := json.Marshal(stu)
	switch kind {
	case reflect.Struct:
		fmt.Println("是结构体")
	}
	if err != nil {
		fmt.Println("serialize failed.ERR=", err)
	}
	fmt.Println(string(marshal))
	return string(marshal)
}
