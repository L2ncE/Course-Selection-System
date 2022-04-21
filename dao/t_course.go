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

func DeleteTCourse(id int) error {
	var Course []model.TCourse
	dbRes := db.Where("Id = ?", id).Delete(&Course)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateTCourseTime(id int, time string) error {
	deRes := db.Model(&model.TCourse{}).Where("id = ?", id).Update("Time", time)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateTCourseTotal(id int, total int) error {
	deRes := db.Model(&model.TCourse{}).Where("id = ?", id).Update("Total", total)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectTCourse() ([]model.TCourse, error) {
	var Course []model.TCourse
	dbRes := db.Model(&model.TCourse{}).Preload("TeacherInfo").Preload("CourseInfo").Find(&Course)
	err := dbRes.Error
	return Course, err
}

func SelectTCourseTime(time string) string {
	course := model.TCourse{}
	dbRes := db.Model(&model.TCourse{}).Where("Time Like ?", "%"+time+"%").Find(&course)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return course.Time
	}
	return course.Time
}

func SelectTCourseTimeById(id int) string {
	var course model.TCourse
	err := db.Model(&model.TCourse{}).Select("time").Where("id = ?", id).First(&course).Error
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return course.Time
	}
	return course.Time
}

func SelectCourseNumById(id int) int {
	var course model.TCourse
	err := db.Model(&model.TCourse{}).Select("CourseNum").Where("id = ?", id).First(&course).Error
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return course.CourseNum
	}
	return course.CourseNum
}
