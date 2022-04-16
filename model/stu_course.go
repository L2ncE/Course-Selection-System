package model

type StuCourse struct {
	Id         int
	StudentNum int `gorm:"column:StudentNum"`
	TCourseNum int `gorm:"column:TCourseNum"`
	Grade      int
}

type StudentNum struct {
	StudentNum int `gorm:"column:StudentNum"`
}
