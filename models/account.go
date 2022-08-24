package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID       int       `gorm:"primary_key" json:"id""`
	Username string    `gorm:"unique"json:"username"`
	Password string    `json:"password"`
	Nickname string    `json:"nickname"`
	Avatar   string    `json:"avatar"`
	Follow   int       `json:"follow"`
	CreateAt time.Time `gorm:"autoCreateTime"json:"createAt"`
	UpdateAt time.Time `gorm:"autoUpdateTime"json:"updateAt"`
	Deleted  int       `default:"0"json:"deleted"`
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

func Info(ID int) (Account, error) {
	account := Account{
		ID:       0,
		Username: "",
		Password: "",
		Nickname: "",
		Avatar:   "",
		Follow:   0,
	}
	err := db.Select("id, nickname, avatar, follow").Where(Account{ID: ID}).First(&account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return account, err
	}

	if account.ID > 0 {
		return account, nil
	}

	return account, nil
}
