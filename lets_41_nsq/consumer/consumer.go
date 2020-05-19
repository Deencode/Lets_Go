// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/19 - 10:18 下午

package main

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

var (
	topic   = "topic_01" // 订阅消息主题名字
	channel = "first"
	// 第一个通道取值
	addr     = "127.0.0.1:4161" // lookupd地址
	consumer *nsq.Consumer
)

type student struct {
	Types           interface{}
	TypesStructName string
	Id              int
	Name            string
	Age             int
}

type MyHandle struct {
	Title string //订阅标题 title 一定要大小
}

func (m *MyHandle) HandleMessage(message *nsq.Message) error {
	// 成功反射出来了
	i := get(message.Body, "student")
	s := i.(*student) //指针
	fmt.Println(s)
	fmt.Println("标题:", m.Title, "NSQD地址:", message.NSQDAddress, "内容:", string(message.Body))
	return nil
}

func get(data []byte, key string) interface{} {
	m := make(map[string]interface{}, 10)
	json.Unmarshal(data, &m) // 反序列化成一个对象
	for n, v := range m {
		switch n {
		// 必须让第三方使用者实行
		case "TypesStructName":
			if v == "student" {
				return &student{
					reflect.Struct,
					func() string {
						return m["TypesStructName"].(string)
					}(),
					int(m["Id"].(float64)),
					string(m["Name"].(string)),
					int(m["Age"].(float64)),
				}
			}
		}
	}

	return nil
}
func init() {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second // 10秒就回调一下
	newConsumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println("create consumer failed.ERR=", err)
		return
	}
	newConsumer.AddHandler(&MyHandle{"消费者consumer"})
	consumer = newConsumer
}
func main() {
	// 开始消费数据 来接lookupd集群
	if err := consumer.ConnectToNSQLookupd(addr); err != nil {
		fmt.Println(err)
	}
	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 如果没有值就会一直阻塞
}
