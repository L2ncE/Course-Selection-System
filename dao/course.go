package dao

import (
	"CSA/model"
	"fmt"
)

//对课程的增删改查

func InsertCourse(course model.Course) error {
	deres := db.Select("Name", "Credit", "Time", "Total").Create(&model.Course{Name: course.Name, Credit: course.Credit, Time: course.Time, Total: course.Total})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteCourse(id int) error {
	var Course []model.Course
	dbRes := db.Where("Id = ?", id).Delete(&Course)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateCourseTime(id int, time string) error {
	deRes := db.Model(&model.Course{}).Where("id = ?", id).Update("Time", time)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateCourseCredit(id int, credit string) error {
	deRes := db.Model(&model.Course{}).Where("id = ?", id).Update("Credit", credit)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateCourseName(id int, name string) error {
	deRes := db.Model(&model.Course{}).Where("id = ?", id).Update("Name", name)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateCourseTotal(id int, total int) error {
	deRes := db.Model(&model.Course{}).Where("id = ?", id).Update("Total", total)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
