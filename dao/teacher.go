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

func UpdateTPassword(Name string, newPassword string) error {
	deRes := db.Model(&model.Teacher{}).Where("Name = ?", Name).Update("Password", newPassword)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectTInfoByName(Name string) (model.TeacherInfo, error) {
	var user model.TeacherInfo
	dbRes := db.Model(&model.Teacher{}).Select("id", "password").Where("Name = ?", Name).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
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

func SelectIdByTInfo(user model.Teacher) int {
	users := model.TeacherInfo{}
	db.Model(&model.Teacher{}).Select("id").Where("Name = ? AND Password = ? AND Question = ? AND Answer = ?", user.Name, user.Password, user.Question, user.Answer).Find(&users)
	return users.Id
}
