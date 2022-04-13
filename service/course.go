package service

import (
	"CSA/dao"
	"CSA/model"
)

func RegisterCourse(c model.Course) error {
	err := dao.InsertCourse(c)
	return err
}

func RemoveCourse(id int) error {
	err := dao.DeleteCourse(id)
	return err
}

//func ChangeCourseTime(id int, time string) error {
//	err := dao.UpdateCourseTime(id, time)
//	return err
//}

func ChangeCourseCredit(id int, credit string) error {
	err := dao.UpdateCourseCredit(id, credit)
	return err
}

func ChangeCourseName(id int, name string) error {
	err := dao.UpdateCourseName(id, name)
	return err
}

//func ChangeCourseTotal(id int, total int) error {
//	err := dao.UpdateCourseTotal(id, total)
//	return err
//}

func GetAllCourse() ([]model.Course, error) {
	course, err := dao.SelectCourse()
	return course, err
}
