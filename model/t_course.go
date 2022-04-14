package model

type TCourse struct {
	Id         int
	CourseNum  int `gorm:"column:CourseNum"`
	TeacherNum int `gorm:"column:TeacherNum"`
	Time       int
	Num        int
	Total      int
}
