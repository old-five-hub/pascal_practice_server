package models

import "gorm.io/gorm"

type Account struct {
	ID       int    `gorm:"primary_key" json:"id""`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(username, password string) (bool, error) {
	var account Account
	err := db.Select("id").Where(Account{Username: username, Password: password}).First(&account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if account.ID > 0 {
		return true, nil
	}

	return false, nil
}
