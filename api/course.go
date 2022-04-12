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

func deleteCourse(ctx *gin.Context) {
	JudgeAdmin(ctx)
	Sid := ctx.Param("id")
	id, _ := strconv.Atoi(Sid)
	err := service.RemoveCourse(id)
	if err != nil {
		fmt.Println("delete err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func UpdateCourseTime(ctx *gin.Context) {
	JudgeAdmin(ctx)
	Sid := ctx.Param("id")
	time := ctx.PostForm("time")
	id, _ := strconv.Atoi(Sid)
	err := service.ChangeCourseTime(id, time)
	if err != nil {
		fmt.Println("update err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func UpdateCourseCredit(ctx *gin.Context) {
	JudgeAdmin(ctx)
	Sid := ctx.Param("id")
	credit := ctx.PostForm("credit")
	id, _ := strconv.Atoi(Sid)
	err := service.ChangeCourseCredit(id, credit)
	if err != nil {
		fmt.Println("update err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func UpdateCourseName(ctx *gin.Context) {
	JudgeAdmin(ctx)
	Sid := ctx.Param("id")
	name := ctx.PostForm("name")
	id, _ := strconv.Atoi(Sid)
	err := service.ChangeCourseName(id, name)
	if err != nil {
		fmt.Println("update err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func UpdateCourseTotal(ctx *gin.Context) {
	JudgeAdmin(ctx)
	Sid := ctx.Param("id")
	Stotal := ctx.PostForm("total")
	id, _ := strconv.Atoi(Sid)
	total, _ := strconv.Atoi(Stotal)
	err := service.ChangeCourseTotal(id, total)
	if err != nil {
		fmt.Println("update err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func getAllCourse(ctx *gin.Context) {
	err := service.GetAllCourse()
	if err != nil {
		fmt.Println("get err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}
