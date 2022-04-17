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

func GetAllStuCourse(id int) ([]model.StuCourseInfo, error) {
	course, err := dao.SelectStuCourse(id)
	return course, err
}
