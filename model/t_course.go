package model

type TCourse struct {
	Id          int
	CourseNum   int `gorm:"column:CourseNum"`
	TeacherNum  int `gorm:"column:TeacherNum"`
	Time        string
	Num         int
	Total       int
	CourseInfo  Course       `gorm:"foreignKey:Id;references:CourseNum"`
	TeacherInfo TeacherInfo2 `gorm:"foreignKey:Id;references:TeacherNum"`
}

type TCourseInfo struct {
	Id          int
	CourseNum   int `gorm:"column:CourseNum"`
	TeacherNum  int `gorm:"column:TeacherNum"`
	Num         int
	CourseInfo  Course       `gorm:"foreignKey:Id;references:CourseNum"`
	TeacherInfo TeacherInfo2 `gorm:"foreignKey:Id;references:TeacherNum"`
}

type SearchTCourse struct {
	Id          int
	CourseNum   int `gorm:"column:CourseNum"`
	TeacherNum  int `gorm:"column:TeacherNum"`
	Time        string
	Num         int
	Total       int
	TeacherInfo TeacherInfo2 `gorm:"foreignKey:Id;references:TeacherNum"`
}

func (TCourseInfo) TableName() string {
	return "t_course"
}

func (SearchTCourse) TableName() string {
	return "t_course"
}
