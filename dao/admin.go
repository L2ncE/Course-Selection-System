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
