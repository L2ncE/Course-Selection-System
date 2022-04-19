package model

type StuCourse struct {
	Id         int
	StudentNum int `gorm:"column:StudentNum"`
	TCourseNum int `gorm:"column:TCourseNum"`
	Grade      int
	Time       string
}

type StuTwoT struct {
	StudentNum int             `gorm:"column:StudentNum"`
	Students   []Student2TInfo `gorm:"foreignKey:Id;references:StudentNum"`
}

type StuCourseInfo struct {
	TCourseNum  int `gorm:"column:TCourseNum"`
	Grade       string
	TCourseInfo []TCourseInfo `gorm:"foreignKey:Id;references:TCourseNum"`
}

func (StuTwoT) TableName() string {
	return "stu_course"
}

func (StuCourseInfo) TableName() string {
	return "stu_course"
}
