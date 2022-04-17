package api

import (
	"CSA/model"
	"CSA/service"
	"CSA/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func registerStuCourse(ctx *gin.Context) {
	IStudentNum, _ := ctx.Get("id")
	studentNum := IStudentNum.(int)

	SCourseNum := ctx.Param("course_num")

	if SCourseNum != "" {
		TCourseNum, _ := strconv.Atoi(SCourseNum)

		course := model.StuCourse{
			StudentNum: studentNum,
			TCourseNum: TCourseNum,
		}

		err := service.RegisterStuCourse(course)
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

func deleteStuCourse(ctx *gin.Context) {
	Sid := ctx.Param("id")
	id, _ := strconv.Atoi(Sid)
	TCourseNum := service.GetTCourseNumByStuCourseId(id)
	err := service.RemoveStuCourse(id, TCourseNum)
	if err != nil {
		fmt.Println("delete err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
	return
}

func getAllStuCourse(ctx *gin.Context) {
	Iid, _ := ctx.Get("id")
	id, _ := Iid.(int)
	course, err := service.GetAllStuCourse(id)
	name := service.GetNameByStuId(id)
	if err != nil {
		fmt.Println("get err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	data1 := "你好" + name
	tool.RespSuccessfulWithTwoData(ctx, data1, course)
	return
}
