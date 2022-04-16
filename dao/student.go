package dao

import (
	"CSA/model"
	"fmt"
)

func InsertStu(user model.Student) error {
	deres := db.Select("Name", "Password", "Question", "Answer", "MajorNum", "NickName").Create(&model.Student{Name: user.Name, Password: user.Password, Question: user.Question, Answer: user.Answer, MajorNum: user.MajorNum, NickName: user.NickName})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateStuPassword(id int, newPassword string) error {
	deRes := db.Model(&model.Student{}).Where("Id = ?", id).Update("Password", newPassword)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectStuInfoById(id int) (model.StudentInfo, error) {
	var user model.StudentInfo
	dbRes := db.Model(&model.Student{}).Select("name", "nickname", "password").Where("Id = ?", id).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func SelectAnswerByStuId(id int) string {
	user := model.Student{}
	db.Model(&model.Student{}).Select("answer").Where("Id = ?", id).Find(&user)
	return user.Answer
}

func SelectQuestionByStuId(id int) string {
	user := model.Student{}
	db.Model(&model.Student{}).Select("question").Where("Id = ?", id).Find(&user)
	return user.Question
}

func SelectNameByStuId(id int) string {
	user := model.Student{}
	db.Model(&model.Student{}).Select("name").Where("Id = ?", id).Find(&user)
	return user.Name
}

func SelectIdByStuNickName(name string) int {
	user := model.Student{}
	db.Model(&model.Student{}).Select("id").Where("nickname = ?", name).Find(&user)
	return user.Id
}

func SelectMajorByStuId(id int) int {
	user := model.Student{}
	db.Model(&model.Student{}).Select("MajorNum").Where("Id = ?", id).Find(&user)
	return user.MajorNum
}

func SelectUserByStuNickName(Name string) (model.StudentInfo, error) {
	var user model.StudentInfo
	dbRes := db.Model(&model.Student{}).Select("id", "password").Where("NickName = ?", Name).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}
