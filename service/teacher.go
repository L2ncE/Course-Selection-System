package service

import (
	"CSA/dao"
	"CSA/model"
	"gorm.io/gorm"
)

func RegisterT(user model.Teacher) error {
	err := dao.InsertT(user)
	return err
}

func ChangeTPassword(id int, newPassword string) error {
	err := dao.UpdateTPassword(id, newPassword)
	return err
}

func IsTPasswordCorrect(id int, password string) (bool, error) {
	user, err := dao.SelectTInfoById(id)
	if err != nil {
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

func GetAnswerByTId(id int) string {
	answer := dao.SelectAnswerByTId(id)
	return answer
}

func GetQuestionByTId(id int) string {
	question := dao.SelectQuestionByTId(id)
	return question
}

//func GetIDByTInfoName(teacher model.Teacher) int {
//	id := dao.SelectIdByTInfo(teacher)
//	return id
//}

func GetNameByTId(id int) string {
	name := dao.SelectNameByTId(id)
	return name
}

func GetIdByTNickName(name string) int {
	id := dao.SelectIdByTNickName(name)
	return id
}

func IsRepeatTNickName(username string) (bool, error) {
	_, err := dao.SelectUserByTNickName(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
