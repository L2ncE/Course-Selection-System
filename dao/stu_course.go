package dao

import (
	"CSA/model"
	"fmt"
	"gorm.io/gorm"
)

func InsertStuCourse(course model.StuCourse) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Select("StudentNum", "TCourseNum").Create(&model.StuCourse{StudentNum: course.StudentNum, TCourseNum: course.TCourseNum}).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.TCourse{}).Where("id = ?", course.TCourseNum).Update("Num", gorm.Expr("Num + 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteStuCourse(id, TCourseNum int) error {
	var Course []model.StuCourse
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("Id = ?", id).Delete(&Course).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.TCourse{}).Where("id = ?", TCourseNum).Update("Num", gorm.Expr("Num - 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectStuCourse(id int) ([]model.StuCourseInfo, error) {
	var Course []model.StuCourseInfo
	dbRes := db.Debug().Where("StudentNum = ?", id).Preload("TCourseInfo").Preload("TCourseInfo.CourseInfo").Preload("TCourseInfo.TeacherInfo").Find(&Course)
	err := dbRes.Error
	return Course, err
}

func SelectTCourseNumByStuCourseId(id int) int {
	course := model.StuCourse{}
	db.Model(&model.StuCourse{}).Select("TCourseNum").Where("Id = ?", id).Find(&course)
	return course.TCourseNum
}
