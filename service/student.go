package service

import (
	"CSA/dao"
	"CSA/model"
	"gorm.io/gorm"
)

func RegisterStu(user model.Student) error {
	err := dao.InsertStu(user)
	return err
}

func ChangeStuPassword(id int, newPassword string) error {
	err := dao.UpdateStuPassword(id, newPassword)
	return err
}

func IsStuPasswordCorrect(id int, password string) (bool, error) {
	user, err := dao.SelectStuInfoById(id)
	if err != nil {
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

func GetAnswerByStuId(id int) string {
	answer := dao.SelectAnswerByStuId(id)
	return answer
}

func GetQuestionByStuId(id int) string {
	question := dao.SelectQuestionByStuId(id)
	return question
}

func GetNameByStuId(id int) string {
	name := dao.SelectNameByStuId(id)
	return name
}

func GetIdByStuNickName(name string) int {
	id := dao.SelectIdByStuNickName(name)
	return id
}

func IsRepeatStuNickName(username string) (bool, error) {
	_, err := dao.SelectUserByStuNickName(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
