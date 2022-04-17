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
	time := ctx.PostForm("time")
	STotal := ctx.PostForm("total")

	if SCourseNum != "" && time != "" && STotal != "" {
		courseNum, _ := strconv.Atoi(SCourseNum)
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

func deleteTCourse(ctx *gin.Context) {
	if JudgeTeacher(ctx) == false {
		return
	}
	Sid := ctx.Param("id")
	id, _ := strconv.Atoi(Sid)
	err := service.RemoveTCourse(id)
	if err != nil {
		fmt.Println("delete err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func UpdateTCourseTime(ctx *gin.Context) {
	if JudgeTeacher(ctx) == false {
		return
	}
	Sid := ctx.Param("id")
	time := ctx.PostForm("time")
	id, _ := strconv.Atoi(Sid)
	err := service.ChangeTCourseTime(id, time)
	if err != nil {
		fmt.Println("update err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func UpdateTCourseTotal(ctx *gin.Context) {
	if JudgeTeacher(ctx) == false {
		return
	}
	Sid := ctx.Param("id")
	STotal := ctx.PostForm("total")
	id, _ := strconv.Atoi(Sid)
	total, _ := strconv.Atoi(STotal)
	err := service.ChangeTCourseTotal(id, total)
	if err != nil {
		fmt.Println("update err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func getAllTCourse(ctx *gin.Context) {
	course, err := service.GetAllTCourse()
	if err != nil {
		fmt.Println("get err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, course)
	return
}
