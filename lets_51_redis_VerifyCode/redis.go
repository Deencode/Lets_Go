// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/8/20 - 5:35 PM

package main

import (
	"fmt"
	redis "github.com/go-redis/redis"
)

// Rdb is redis client pointer
var Rdb *redis.Client

const (
	host     = "128.199.155.162"
	port     = "6379"
	password = "deen.job"
)

func init() {
	// init redis main connection
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
	})
	Rdb = client
}
