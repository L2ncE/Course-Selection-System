package model

type StuCourse struct {
	Id         int
	StudentNum int `gorm:"column:StudentNum"`
	TCourseNum int `gorm:"column:TCourseNum"`
	Grade      int
}

type StuTwoT struct {
	StudentNum int             `gorm:"column:StudentNum"`
	Students   []Student2TInfo `gorm:"foreignKey:Id;references:StudentNum"`
}

func (StuTwoT) TableName() string {
	return "stu_course"
}
