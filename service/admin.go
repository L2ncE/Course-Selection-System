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

func IsAdmin(id int) (bool, error) {
	identity, err := dao.SelectIdentityById(id)
	if err != nil {
		return false, err
	}
	if identity != 3 {
		return false, nil
	}
	return true, nil
}

func GetIdByAdminUserName(name string) int {
	id := dao.SelectIdByAdminUserName(name)
	return id
}
