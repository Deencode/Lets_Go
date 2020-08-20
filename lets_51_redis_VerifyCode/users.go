// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/8/20 - 5:50 下午 - UTC/GMT+08:00

package main

var UserList map[int64]*User

type User struct {
	Account  string `json:"account"'`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

func init() {
	UserList = make(map[int64]*User, 5)
	UserList[1] = &User{"admin1", "123456", "管理员1"}
	UserList[2] = &User{"admin2", "123456", "管理员2"}
	UserList[3] = &User{"admin3", "123456", "管理员3"}
}

type UserMgr struct{}

func NewUserMgr() *UserMgr {
	return new(UserMgr)
}

func (m *UserMgr) Verify(account, password string) bool {
	for _, user := range UserList {
		if user.Account == account && user.Password == password {
			return true
		}
	}
	return false
}
