package service

import (
	"CSA/dao"
	"CSA/model"
)

func RegisterStuCourse(c model.StuCourse) error {
	err := dao.InsertStuCourse(c)
	return err
}

func RemoveStuCourse(id int) error {
	err := dao.DeleteStuCourse(id)
	return err
}

func GetAllStuCourse() ([]model.StuCourse, error) {
	course, err := dao.SelectStuCourse()
	return course, err
}
