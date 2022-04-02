package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(CORS())

	err := engine.Run(":8081")
	if err != nil {
		fmt.Printf("init error:%v\n", err)
		return
	}
}
