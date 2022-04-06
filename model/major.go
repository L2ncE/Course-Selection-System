package model

type Major struct {
	MajorNum  string `gorm:"column:MajorNum"`
	MajorName string `gorm:"column:MajorName"`
}
