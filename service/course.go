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

func ChangeCourseCredit(id int, credit float64) error {
	err := dao.UpdateCourseCredit(id, credit)
	return err
}

func ChangeCourseName(id int, name string) error {
	err := dao.UpdateCourseName(id, name)
	return err
}

func GetAllCourse() ([]model.Course, error) {
	course, err := dao.SelectCourse()
	return course, err
}

func GetCreditById(id int) float64 {
	credit := dao.SelectCreditByCourseNum(id)
	return credit
}

func GetSearch(str string) ([]model.SearchCourse, error) {
	course, err := dao.Search(str)
	return course, err
}
