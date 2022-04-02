package service

import (
	"CSA/dao"
	"CSA/model"
)

func RegisterT(user model.Teacher) error {
	err := dao.InsertT(user)
	return err
}

func GetAnswerByTName(Name string) string {
	answer := dao.SelectAnswerByTName(Name)
	return answer
}

func GetQuestionByTName(Name string) string {
	question := dao.SelectQuestionByTName(Name)
	return question
}
