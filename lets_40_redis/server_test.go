// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/18 - 5:40 下午

package main

import "testing"

func TestRedis(t *testing.T) {
	//initClient()
	t.Log(initClient())
	t.Log(Rdb.Get("name"))
}
