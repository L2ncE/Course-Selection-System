package api

import (
	"CSA/model"
	"CSA/service"
	"CSA/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func JudgeAdmin(ctx *gin.Context) {
	Iid, _ := ctx.Get("id")
	id := Iid.(int)
	flag, err := service.IsAdmin(id)
	if err != nil {
		fmt.Println("judge err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithData(ctx, "你不是管理员，无法操作")
		return
	}
}
func registerCourse(ctx *gin.Context) {
	JudgeAdmin(ctx)
	name := ctx.PostForm("name")
	credit := ctx.PostForm("credit")
	time := ctx.PostForm("time")
	total := ctx.PostForm("total")

	fCredit, _ := strconv.ParseFloat(credit, 32)
	iTotal, _ := strconv.Atoi(total)
	if name != "" && credit != "" && time != "" && total != "" {
		l1 := len([]rune(name))
		l2 := len([]rune(total))
		if l1 > 20 || l1 < 1 {
			tool.RespErrorWithData(ctx, "课程名请在1到20位之间")
			return
		}
		if l2 > 500 || l2 <= 0 {
			tool.RespErrorWithData(ctx, "总人数请设置在0-500")
			return
		}
		if fCredit > 20 || fCredit <= 0 {
			tool.RespErrorWithData(ctx, "学分请在0-20之间")
			return
		}

		course := model.Course{
			Name:   name,
			Credit: fCredit,
			Time:   time,
			Total:  iTotal,
		}

		err := service.RegisterCourse(course)
		if err != nil {
			fmt.Println("register err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
		return
	}
	tool.RespErrorWithData(ctx, "请将信息输入完整")
	return
}
