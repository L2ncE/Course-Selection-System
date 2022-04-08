package dao

import (
	"CSA/model"
	"fmt"
)

func InsertT(user model.Teacher) error {
	deres := db.Select("Name", "NickName", "Password", "Question", "Answer").Create(&model.Teacher{Name: user.Name, NickName: user.NickName, Password: user.Password, Question: user.Question, Answer: user.Answer})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateTPassword(id int, newPassword string) error {
	deRes := db.Model(&model.Teacher{}).Where("Id = ?", id).Update("Password", newPassword)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectTInfoById(id int) (model.TeacherInfo, error) {
	var user model.TeacherInfo
	dbRes := db.Model(&model.Teacher{}).Select("name", "nickname", "password").Where("Id = ?", id).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func SelectAnswerByTId(id int) string {
	user := model.Teacher{}
	db.Model(&model.Teacher{}).Select("answer").Where("Id = ?", id).Find(&user)
	return user.Answer
}

func SelectQuestionByTId(id int) string {
	user := model.Teacher{}
	db.Model(&model.Teacher{}).Select("question").Where("Id = ?", id).Find(&user)
	return user.Question
}

//func SelectIdByTInfo(user model.Teacher) int {
//	users := model.TeacherInfo{}
//	db.Model(&model.Teacher{}).Select("id").Where("Name = ? AND Password = ? AND Question = ? AND Answer = ?", user.Name, user.Password, user.Question, user.Answer).Find(&users)
//	return users.Id
//}

func SelectNameByTId(id int) string {
	user := model.Teacher{}
	db.Model(&model.Teacher{}).Select("name").Where("Id = ?", id).Find(&user)
	return user.Name
}

func SelectIdByTNickName(name string) int {
	user := model.Teacher{}
	db.Model(&model.Teacher{}).Select("id").Where("nickname = ?", name).Find(&user)
	return user.Id
}
