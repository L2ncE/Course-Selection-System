package dao

import (
	"CSA/model"
	"fmt"
)

//对课程的增删改查

func InsertCourse(course model.Course) error {
	deres := db.Select("Name", "Credit").Create(&model.Course{Name: course.Name, Credit: course.Credit})
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

func UpdateCourseCredit(id int, credit float64) error {
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

func SelectCourse() ([]model.Course, error) {
	var Course []model.Course
	dbRes := db.Model(&model.Course{}).Find(&Course)
	err := dbRes.Error
	return Course, err
}

func SelectCreditByCourseNum(id int) float64 {
	course := model.Course{}
	db.Model(&model.Course{}).Select("Credit").Where("Id = ?", id).Find(&course)
	return course.Credit
}

func Search(str string) ([]model.SearchCourse, error) {
	fmt.Print(str)
	var Course []model.SearchCourse
	dbRes := db.Where("Name LIKE ?", str).Preload("TCourse").Preload("TCourse.TeacherInfo").Find(&Course)
	err := dbRes.Error
	return Course, err
}
