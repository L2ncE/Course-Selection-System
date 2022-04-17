package dao

import (
	"CSA/model"
	"fmt"
	"gorm.io/gorm/clause"
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

func SelectUserByTNickName(Name string) (model.TeacherInfo, error) {
	var user model.TeacherInfo
	dbRes := db.Model(&model.Teacher{}).Select("id", "password").Where("NickName = ?", Name).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func SelectTeacherIdentityById(id int) (identity int, err error) {
	var user model.Teacher
	db.Model(&model.Teacher{}).Select("identity").Where("id = ?", id).First(&user)
	if err != nil {
		return user.Identity, err
	}
	fmt.Println(user)
	return user.Identity, nil
}

func SelectStudentInfoByTCourseNum(id int) ([]model.StuTwoT, error) {
	var info []model.StuTwoT
	err := db.Debug().Where("TCourseNum = ?", id).Preload(clause.Associations).Find(&info).Error
	if err != nil {
		return info, err
	}
	return info, nil
}
