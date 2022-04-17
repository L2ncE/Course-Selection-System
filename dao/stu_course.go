package dao

import (
	"CSA/model"
	"fmt"
)

func InsertStuCourse(course model.StuCourse) error {
	deres := db.Select("StudentNum", "TCourseNum").Create(&model.StuCourse{StudentNum: course.StudentNum, TCourseNum: course.TCourseNum})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteStuCourse(id int) error {
	var Course []model.StuCourse
	dbRes := db.Where("Id = ?", id).Delete(&Course)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectStuCourse(id int) ([]model.StuCourseInfo, error) {
	var Course []model.StuCourseInfo
	dbRes := db.Model(&model.StuCourse{}).Where("StudentNum = ?", id).Find(&Course)
	err := dbRes.Error
	return Course, err
}
