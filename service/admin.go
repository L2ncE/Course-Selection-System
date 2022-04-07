package service

import "CSA/dao"

func IsAdminPasswordCorrect(Name, password string) (bool, error) {
	user, err := dao.SelectAdminByName(Name)
	if err != nil {
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}
