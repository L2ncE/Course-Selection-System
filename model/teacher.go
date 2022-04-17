package model

type Teacher struct {
	Id       int
	Name     string
	NickName string `gorm:"column:NickName"`
	Password string
	Question string
	Answer   string
	Identity int
}

type TeacherInfo struct {
	Id       int
	Name     string
	NickName string `gorm:"column:NickName"`
	Password string
}

type TeacherInfo2 struct {
	Id   int
	Name string
}

func (TeacherInfo2) TableName() string {
	return "teacher"
}
