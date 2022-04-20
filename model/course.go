package model

type Course struct {
	Id     int
	Name   string
	Credit float64
}

type SearchCourse struct {
	Id      int
	Name    string
	Credit  float64
	TCourse []SearchTCourse `gorm:"foreignKey:CourseNum;references:Id"`
}

func (SearchCourse) TableName() string {
	return "course"
}
