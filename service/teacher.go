package service

import (
	"CSA/dao"
	"CSA/model"
)

func RegisterT(user model.Teacher) error {
	err := dao.InsertT(user)
	return err
}

func ChangeTPassword(Name, newPassword string) error {
	err := dao.UpdateTPassword(Name, newPassword)
	return err
}

func IsTPasswordCorrect(Name, password string) (bool, error) {
	user, err := dao.SelectTInfoByName(Name)
	if err != nil {
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

func GetAnswerByTName(Name string) string {
	answer := dao.SelectAnswerByTName(Name)
	return answer
}

func GetQuestionByTName(Name string) string {
	question := dao.SelectQuestionByTName(Name)
	return question
}
