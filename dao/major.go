package dao

import "CSA/model"

func SelectAllMajor() ([]model.Major, error) {
	var majors []model.Major

	dbRes := db.Model(&model.Major{}).Find(&majors)
	if dbRes.Error != nil {
		return majors, dbRes.Error
	}
	return majors, nil
}

func SelectMajorNameByMajorId(id int) string {
	major := model.Major{}
	db.Model(&model.Major{}).Select("MajorName").Where("MajorNum = ?", id).Find(&major)
	return major.MajorName
}
