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
