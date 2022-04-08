package model

type Teacher struct {
	Id       int
	Name     string
	NickName string `gorm:"column:NickName"`
	Password string
	Question string
	Answer   string
}

type TeacherInfo struct {
	Id       int
	Name     string
	NickName string `gorm:"column:NickName"`
	Password string
}
