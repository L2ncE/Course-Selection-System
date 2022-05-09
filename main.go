package main

import (
	"CSA/api"
	"CSA/config"
	"CSA/dao"
	"CSA/pprof"
	"fmt"
)

func main() {
	pprof.InitPprofMonitor()
	config.InitConfig()
	if err := dao.InitGormDB(); err != nil {
		fmt.Printf("init gorm failed, err:%v\n", err)
		return
	}
	fmt.Println("连接GORM MySQL数据库成功!")
	api.InitEngine()
}
