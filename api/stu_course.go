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
	TCourseNum, _ := strconv.Atoi(SCourseNum)
	time := service.GetTCourseTimeById(TCourseNum)
	flag := service.IsRepeatTime(time)

	if flag {
		tool.RespErrorWithData(ctx, "选课冲突")
		return
	}

	if SCourseNum != "" {

		course := model.StuCourse{
			Time:       time,
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
	SCourseId := ctx.Param("id")
	Iid, _ := ctx.Get("id")
	id := Iid.(int)
	CourseId, _ := strconv.Atoi(SCourseId)

	StudentNum := service.GetStudentNumByStuCourseId(CourseId)
	TCourseNum := service.GetTCourseNumByStuCourseId(CourseId)
	fmt.Println(StudentNum, id)
	if StudentNum != id {
		tool.RespErrorWithData(ctx, "您不能删除他人的课程")
		return
	}
	err := service.RemoveStuCourse(CourseId, TCourseNum)
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
