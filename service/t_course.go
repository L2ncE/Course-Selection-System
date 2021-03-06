package service

import (
	"CSA/dao"
	"CSA/model"
	"fmt"
)

func RegisterTCourse(c model.TCourse) error {
	err := dao.InsertTCourse(c)
	return err
}

func RemoveTCourse(id int) error {
	err := dao.DeleteTCourse(id)
	return err
}

func ChangeTCourseTime(id int, time string) error {
	err := dao.UpdateTCourseTime(id, time)
	return err
}

func ChangeTCourseTotal(id int, total int) error {
	err := dao.UpdateTCourseTotal(id, total)
	return err
}

func GetAllTCourse() ([]model.TCourse, error) {
	course, err := dao.SelectTCourse()
	return course, err
}

func IsRepeatTime(time string) bool {
	str := dao.SelectTCourseTime(time)
	if str != "" && str != time {
		fmt.Println(str)
		return true
	}
	return false
}

func GetTCourseTimeById(id int) string {
	time := dao.SelectTCourseTimeById(id)
	return time
}
