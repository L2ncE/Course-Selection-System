package api

import (
	"CSA/service"
	"CSA/tool"
	"github.com/gin-gonic/gin"
)

func searchMajor(ctx *gin.Context) {
	res, _ := service.GetAllMajor()
	tool.RespErrorWithData(ctx, res)
}
