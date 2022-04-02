package model

type Student struct {
	Id       int
	Name     string
	Password string
	Question string
	Answer   string
	MajorNum int `gorm:"column:MajorNum"`
}

type StudentInfo struct {
	Id       int
	Password string
}
