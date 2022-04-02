package service

import (
	"CSA/dao"
	"CSA/model"
)

func RegisterStu(user model.Student) error {
	err := dao.InsertStu(user)
	return err
}

func GetAnswerByStuName(Name string) string {
	answer := dao.SelectAnswerByStuName(Name)
	return answer
}

func GetQuestionByStuName(Name string) string {
	question := dao.SelectQuestionByStuName(Name)
	return question
}
