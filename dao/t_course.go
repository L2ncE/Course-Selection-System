package dao

import (
	"CSA/model"
	"fmt"
)

func InsertTCourse(course model.TCourse) error {
	deres := db.Select("CourseNum", "TeacherNum", "Time", "Total").Create(&model.TCourse{CourseNum: course.CourseNum, TeacherNum: course.TeacherNum, Time: course.Time, Total: course.Total})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}
