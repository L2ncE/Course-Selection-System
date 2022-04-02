package main

import (
	"CSA/api"
	"CSA/dao"
	"fmt"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	} else {
		fmt.Println("连接MySQL数据库成功!")
	}
	err = dao.InitRedis()
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
	} else {
		fmt.Println("连接Redis数据库成功!")
	}
	err = dao.InitGormDB()
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
	} else {
		fmt.Println("连接GORM MySQL数据库成功!")
	}
	api.InitEngine()
}
