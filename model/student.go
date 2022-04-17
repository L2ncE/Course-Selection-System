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

type Student2TInfo struct {
	Id       int
	Name     string
	MajorNum int `gorm:"column:MajorNum"`
}

func (Student2TInfo) TableName() string {
	return "student"
}
