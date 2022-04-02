package dao

import (
	"CSA/model"
	"fmt"
)

func InsertT(user model.Teacher) error {
	deres := db.Select("Name", "Password", "Question", "Answer").Create(&model.Teacher{Name: user.Name, Password: user.Password, Question: user.Question, Answer: user.Answer})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectAnswerByTName(Name string) string {
	user := model.Teacher{}
	db.Model(&model.Teacher{}).Select("answer").Where("Name = ?", Name).Find(&user)
	return user.Answer
}

func SelectQuestionByTName(Name string) string {
	user := model.Teacher{}
	db.Model(&model.Teacher{}).Select("question").Where("Name = ?", Name).Find(&user)
	return user.Question
}
