package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(CORS())

	engine.GET("/major", searchMajor)   //得到专业ID NAME 对照表
	engine.GET("/course", getAllCourse) //得到所有课程

	studentGroup := engine.Group("/stu") //学生
	{
		studentGroup.POST("/register", registerStu)    //注册
		studentGroup.POST("/login", loginStu)          //登录
		studentGroup.POST("/ss", stuSecretSecurity)    //密保功能
		studentGroup.POST("/ss/question", stuQuestion) //查询密保问题
		{
			studentGroup.Use(JWTAuth)                        //需要token
			studentGroup.PUT("/password", changeStuPassword) //修改密码
		}
	}

	teacherGroup := engine.Group("/t") //老师
	{
		teacherGroup.POST("/register", registerT)    //注册
		teacherGroup.POST("/login", loginT)          //登录
		teacherGroup.POST("/ss", tSecretSecurity)    //密保功能
		teacherGroup.POST("/ss/question", tQuestion) //查询密保问题
		{
			teacherGroup.Use(JWTAuth)                      //需要token
			teacherGroup.PUT("/password", changeTPassword) //修改密码
		}
	}

	adminGroup := engine.Group("/admin")
	{
		adminGroup.POST("/login", loginAdmin) //登录
	}

	courseGroup := engine.Group("/course")
	{
		courseGroup.Use(JWTAuth)
		courseGroup.POST("/register", registerCourse)      //新增注册课程
		courseGroup.DELETE("/:id", deleteCourse)           //删除课程
		courseGroup.PUT("/time/:id", UpdateCourseTime)     //更新课程时间
		courseGroup.PUT("/credit/:id", UpdateCourseCredit) //更新课程学分
		courseGroup.PUT("/name/:id", UpdateCourseName)     //更新课程名字
		courseGroup.PUT("/total/:id", UpdateCourseTotal)   //更新课程总人数
	}
	err := engine.Run(":8081")
	if err != nil {
		fmt.Printf("init error:%v\n", err)
		return
	}
}
