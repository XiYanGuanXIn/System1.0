package test

//package main

import (
	"System1.0/Dao/UserInfoDao"
	"fmt"
)

func main() {
	fmt.Println(UserInfoDao.ChangeUserInfo(3, "hello world,I am OceanCT"))
	fmt.Println(UserInfoDao.FindUserInfoByID(3))
}
