package service

import (
	"CSA/dao"
	"CSA/model"
)

func RegisterStu(user model.Student) error {
	err := dao.InsertStu(user)
	return err
}

func ChangeStuPassword(Name, newPassword string) error {
	err := dao.UpdateStuPassword(Name, newPassword)
	return err
}

func IsStuPasswordCorrect(Name, password string) (bool, error) {
	user, err := dao.SelectStuInfoByName(Name)
	if err != nil {
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

//func IsRepeatUsername(Name string) (bool, error) {
//	_, err := dao.SelectStuInfoByName(Name)
//	if err != nil {
//		if err == gorm.ErrRecordNotFound {
//			return false, nil
//		}
//		return false, err
//	}
//	return true, nil
//}

func GetAnswerByStuName(Name string) string {
	answer := dao.SelectAnswerByStuName(Name)
	return answer
}

func GetQuestionByStuName(Name string) string {
	question := dao.SelectQuestionByStuName(Name)
	return question
}
