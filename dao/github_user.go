package dao

import (
	"CSA/model"
	"fmt"
)

func InsertGitHub(user model.GitHubUser) error {
	deres := db.Select("Name", "MajorNum", "NickName").Create(&model.GitHubUser{Name: user.Name, MajorNum: user.MajorNum, NickName: user.NickName})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectGitHubByNickName(Name string) (model.GitHubUser, error) {
	var user model.GitHubUser
	dbRes := db.Model(&model.GitHubUser{}).Select("id", "name").Where("NickName = ?", Name).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}
