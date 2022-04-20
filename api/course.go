package api

import (
	"CSA/model"
	"CSA/service"
	"CSA/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func JudgeAdmin(ctx *gin.Context) bool {
	Iid, _ := ctx.Get("id")
	id := Iid.(int)
	flag, err := service.IsAdmin(id)
	if err != nil {
		fmt.Println("judge err: ", err)
		tool.RespInternalError(ctx)
		return false
	}
	if !flag {
		tool.RespErrorWithData(ctx, "你不是管理员，无法操作")
		return false
	}
	return true
}

func registerCourse(ctx *gin.Context) {
	if JudgeAdmin(ctx) == false {
		return
	}

	name := ctx.PostForm("name")
	credit := ctx.PostForm("credit")

	fCredit, _ := strconv.ParseFloat(credit, 32)
	if name != "" && credit != "" {
		l1 := len([]rune(name))
		if l1 > 20 || l1 < 1 {
			tool.RespErrorWithData(ctx, "课程名请在1到20位之间")
			return
		}
		if fCredit > 20 || fCredit <= 0 {
			tool.RespErrorWithData(ctx, "学分请在0-20之间")
			return
		}

		course := model.Course{
			Name:   name,
			Credit: fCredit,
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
	if JudgeAdmin(ctx) == false {
		return
	}
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

func UpdateCourseCredit(ctx *gin.Context) {
	if JudgeAdmin(ctx) == false {
		return
	}
	Sid := ctx.Param("id")
	credit := ctx.PostForm("credit")
	fCredit, _ := strconv.ParseFloat(credit, 32)
	id, _ := strconv.Atoi(Sid)
	err := service.ChangeCourseCredit(id, fCredit)
	if err != nil {
		fmt.Println("update err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func UpdateCourseName(ctx *gin.Context) {
	if JudgeAdmin(ctx) == false {
		return
	}
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

func getAllCourse(ctx *gin.Context) {
	course, err := service.GetAllCourse()
	if err != nil {
		fmt.Println("get err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, course)
	return
}

func search(ctx *gin.Context) {
	PostStr := ctx.PostForm("string")
	str := "%" + PostStr + "%"
	course, err := service.GetSearch(str)
	if err != nil {
		fmt.Println("get err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, course)
	return
}
