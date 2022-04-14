package api

import (
	"CSA/model"
	"CSA/service"
	"CSA/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func JudgeTeacher(ctx *gin.Context) bool {
	Iid, _ := ctx.Get("id")
	id := Iid.(int)
	flag, err := service.IsTeacher(id)
	if err != nil {
		fmt.Println("judge err: ", err)
		tool.RespInternalError(ctx)
		return false
	}
	if !flag {
		tool.RespErrorWithData(ctx, "你不是老师或管理员，无法操作")
		return false
	}
	return true
}

func registerTCourse(ctx *gin.Context) {
	if JudgeTeacher(ctx) == false {
		return
	}

	ITeacherNum, _ := ctx.Get("id")
	teacherNum := ITeacherNum.(int)

	SCourseNum := ctx.Param("course_num")
	STime := ctx.PostForm("time")
	STotal := ctx.PostForm("total")

	if SCourseNum != "" && STime != "" && STotal != "" {
		courseNum, _ := strconv.Atoi(SCourseNum)
		time, _ := strconv.Atoi(STime)
		total, _ := strconv.Atoi(STotal)

		if total <= 0 || total > 1000 {
			tool.RespErrorWithData(ctx, "总人数请在1-1000人之间")
			return
		}

		course := model.TCourse{
			CourseNum:  courseNum,
			TeacherNum: teacherNum,
			Time:       time,
			Total:      total,
		}

		err := service.RegisterTCourse(course)
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
