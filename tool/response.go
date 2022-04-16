package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": data,
	})
}

func RespInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器错误",
	})
}

func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
	})
}

func RespSuccessfulWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
		"data": data,
	})
}

func RespSuccessfulWithTwoData(ctx *gin.Context, data1 interface{}, data2 interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info":  "成功",
		"data1": data1,
		"data2": data2,
	})
}
