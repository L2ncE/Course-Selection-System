package model

type Student struct {
	Id       int
	Name     string
	NickName string `gorm:"column:NickName"`
	Password string
	Question string
	Answer   string
	MajorNum int `gorm:"column:MajorNum"`
	Identity int
}

type StudentInfo struct {
	Id       int
	Name     string
	NickName string `gorm:"column:NickName"`
	Password string
}
