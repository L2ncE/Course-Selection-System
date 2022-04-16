package service

import (
	"CSA/dao"
	"CSA/model"
)

func GetAllMajor() ([]model.Major, error) {
	res, _ := dao.SelectAllMajor()
	return res, nil
}

func GetMajorNameByMajorId(id int) string {
	Major := dao.SelectMajorNameByMajorId(id)
	return Major
}
