package dao

import (
	"CSA/model"
	"fmt"
)

func SelectAdminByName(Name string) (model.Admin, error) {
	var user model.Admin
	dbRes := db.Model(&model.Admin{}).Select("id", "password").Where("Name = ?", Name).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func SelectAdminIdentityById(id int) (identity int, err error) {
	var user model.Admin
	db.Model(&model.Admin{}).Select("identity").Where("id = ?", id).First(&user)
	if err != nil {
		return user.Identity, err
	}
	fmt.Println(user)
	return user.Identity, nil
}

func SelectIdByAdminUserName(name string) int {
	user := model.Admin{}
	db.Model(&model.Admin{}).Select("id").Where("name = ?", name).Find(&user)
	return user.Id
}
