package dao

import (
	"CSA/model"
	"fmt"
)

func InsertStu(user model.Student) error {
	deres := db.Select("Name", "Password", "Question", "Answer", "MajorNum").Create(&model.Student{Name: user.Name, Password: user.Password, Question: user.Question, Answer: user.Answer, MajorNum: user.MajorNum})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateStuPassword(Name string, newPassword string) error {
	deRes := db.Model(&model.Student{}).Where("Name = ?", Name).Update("Password", newPassword)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectStuInfoByName(Name string) (model.StudentInfo, error) {
	var user model.StudentInfo
	dbRes := db.Model(&model.Student{}).Select("id", "password").Where("Name = ?", Name).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func SelectAnswerByStuName(Name string) string {
	user := model.Student{}
	db.Model(&model.Student{}).Select("answer").Where("Name = ?", Name).Find(&user)
	return user.Answer
}

func SelectQuestionByStuName(Name string) string {
	user := model.Student{}
	db.Model(&model.Student{}).Select("question").Where("Name = ?", Name).Find(&user)
	return user.Question
}
