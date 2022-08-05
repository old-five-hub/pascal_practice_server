package models

import "gorm.io/gorm"

type Account struct {
	ID       int    `gorm:"primary_key" json:"id""`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Follow   int    `json:"follow"`
}

func Login(username, password string) (Account, error) {
	account := Account{
		ID:       0,
		Username: username,
		Password: password,
		Nickname: "",
		Avatar:   "",
		Follow:   0,
	}
	err := db.Select("id, nickname, avatar, follow").Where(Account{Username: username, Password: password}).First(&account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return account, err
	}

	if account.ID > 0 {
		return account, nil
	}

	return account, nil
}
