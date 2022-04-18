package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(CORS())

	engine.GET("/major", searchMajor)      //得到专业ID NAME 对照表
	engine.GET("/course", getAllCourse)    //得到所有课程
	engine.GET("/t_course", getAllTCourse) //得到所有老师课程

	studentGroup := engine.Group("/stu") //学生
	{
		studentGroup.POST("/register", registerStu)    //注册
		studentGroup.POST("/login", loginStu)          //登录
		studentGroup.POST("/ss", stuSecretSecurity)    //密保功能
		studentGroup.POST("/ss/question", stuQuestion) //查询密保问题
		{
			studentGroup.Use(JWTAuth)                        //需要token
			studentGroup.PUT("/password", changeStuPassword) //修改密码
			studentGroup.GET("/all", getAllStuCourse)        //得到所有已选课程
		}
	}

	teacherGroup := engine.Group("/t") //老师
	{
		teacherGroup.POST("/register", registerT)    //注册
		teacherGroup.POST("/login", loginT)          //登录
		teacherGroup.POST("/ss", tSecretSecurity)    //密保功能
		teacherGroup.POST("/ss/question", tQuestion) //查询密保问题
		{
			teacherGroup.Use(JWTAuth)                                      //需要token
			teacherGroup.PUT("/password", changeTPassword)                 //修改密码
			teacherGroup.GET("/course/:id", searchStudentInfoByTCourseNum) //得到老师自己课程的学生详情
		}
	}

	adminGroup := engine.Group("/admin")
	{
		adminGroup.POST("/login", loginAdmin) //登录
	}

	courseGroup := engine.Group("/course") //课程
	{
		courseGroup.Use(JWTAuth)
		courseGroup.POST("/register", registerCourse)      //新增注册课程
		courseGroup.DELETE("/:id", deleteCourse)           //删除课程
		courseGroup.PUT("/credit/:id", UpdateCourseCredit) //更新课程学分
		courseGroup.PUT("/name/:id", UpdateCourseName)     //更新课程名字
	}

	tCourseGroup := engine.Group("/t_course") //老师课程
	{
		tCourseGroup.Use(JWTAuth)
		tCourseGroup.POST("/register/:course_num", registerTCourse) //新增注册老师课程
		tCourseGroup.DELETE("/:id", deleteTCourse)                  //删除老师课程
		tCourseGroup.PUT("/time/:id", UpdateTCourseTime)            //更新课程时间
		tCourseGroup.PUT("/total/:id", UpdateTCourseTotal)          //更新课程总人数
	}

	stuCourseGroup := engine.Group("/stu_course") //学生课程
	{
		stuCourseGroup.Use(JWTAuth)
		stuCourseGroup.POST("/register/:course_num", registerStuCourse) //新增注册学生课程
		stuCourseGroup.DELETE("/:id", deleteStuCourse)                  //删除学生课程
	}

	gitCourseGroup := engine.Group("/git") //GitHub用户
	{
		gitCourseGroup.POST("/oauth", Oauth)
	}

	err := engine.Run(":8081")
	if err != nil {
		fmt.Printf("init error:%v\n", err)
		return
	}
}
