package service

import (
	"CSA/dao"
	"CSA/model"
)

func RegisterStuCourse(c model.StuCourse) error {
	err := dao.InsertStuCourse(c)
	return err
}

func RemoveStuCourse(id, TCourseNum int) error {
	err := dao.DeleteStuCourse(id, TCourseNum)
	return err
}

func GetAllStuCourse(id int) ([]model.StuCourseInfo, error) {
	course, err := dao.SelectStuCourse(id)
	return course, err
}

func GetTCourseNumByStuCourseId(id int) int {
	TCourseNum := dao.SelectTCourseNumByStuCourseId(id)
	return TCourseNum
}
