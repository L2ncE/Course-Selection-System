package service

import (
	"CSA/dao"
	"CSA/model"
)

func RegisterStuCourse(c model.StuCourse, credit float64) error {
	err := dao.InsertStuCourse(c, credit)
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

func GetStudentNumByStuCourseId(id int) int {
	StudentNum := dao.SelectStudentNumByStuCourseId(id)
	return StudentNum
}

func GetCourseNumById(id int) int {
	CourseNum := dao.SelectCourseNumById(id)
	return CourseNum
}

func ChangeGrade(id int, newGrade string) error {
	err := dao.UpdateGrade(id, newGrade)
	return err
}
