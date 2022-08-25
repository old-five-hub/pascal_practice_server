package models

import (
	"time"
)

const (
	LikeQuestion int = 0
	LikeComment      = 1
)

type UserLike struct {
	ID         int       `gorm:"primary_key" json:"id""`
	Account    Account   `json:"nickName" `
	AccountID  int       `json:"accountId"`
	TypeID     int       `json:"typeId"`
	LikeType   int       `json:"likeType"`
	LikeStatus int       `default:"0"json:"listStatus"`
	CreateAt   time.Time `gorm:"autoCreateTime"json:"createAt"`
	UpdateAt   time.Time `gorm:"autoUpdateTime"json:"updateAt"`
	Deleted    int       `default:"0"json:"deleted"`
}

type UserLikeStat struct {
	ID        int       `gorm:"primary_key" json:"id""`
	TypeID    int       `json:"typeId"`
	LikeType  int       `json:"likeType"`
	LikeCount int       `json:"likeCount"`
	CreateAt  time.Time `gorm:"autoCreateTime"json:"createAt"`
	UpdateAt  time.Time `gorm:"autoUpdateTime"json:"updateAt"`
	Deleted   int       `default:"0"json:"deleted"`
}

func UpdateUserLike(id int, data map[string]interface{}) error {
	likeStatus := data["likeStatus"].(int)
	if err := db.Model(&UserLike{}).Where("id = ?", id).Update("like_status", likeStatus).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserLike(data map[string]interface{}) error {
	userLike := &UserLike{
		AccountID:  data["accountId"].(int),
		TypeID:     data["typeId"].(int),
		LikeType:   data["likeType"].(int),
		LikeStatus: data["likeStatus"].(int),
	}
	if err := db.Create(userLike).Error; err != nil {
		return err
	}
	return nil
}

func ExistUserLike(data map[string]interface{}) (UserLike, error) {
	userLike := &UserLike{
		AccountID: data["accountId"].(int),
		TypeID:    data["typeId"].(int),
		LikeType:  data["likeType"].(int),
	}

	existResult := UserLike{}

	err := db.Where("account_id = ? AND type_id = ? AND like_type = ?", userLike.AccountID, userLike.TypeID, userLike.LikeType).Find(&existResult).Error
	if err != nil {
		return existResult, err
	}

	return existResult, nil
}

func UpdateUserLikeStat(id int, data map[string]interface{}) error {
	if err := db.Model(&UserLikeStat{}).Where("id = ?", id).Update("like_count", data["likeCount"].(int)).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserLikeStat(data map[string]interface{}) error {

	userLikeStat := UserLikeStat{
		TypeID:    data["typeId"].(int),
		LikeType:  data["likeType"].(int),
		LikeCount: data["likeCount"].(int),
	}
	if err := db.Create(&userLikeStat).Error; err != nil {
		return err
	}
	return nil
}

func ExistUserLikeStat(data map[string]interface{}) (UserLikeStat, error) {
	userLikeStat := &UserLikeStat{
		TypeID:   data["typeId"].(int),
		LikeType: data["likeType"].(int),
	}

	existResult := UserLikeStat{}

	err := db.Where("type_id = ? AND like_type = ?", userLikeStat.TypeID, userLikeStat.LikeType).Find(&existResult).Error

	if err != nil {
		return existResult, err
	}

	return existResult, nil
}

func GetUserLikeCount(data map[string]interface{}) (int, error) {
	userLikeStat, err := ExistUserLikeStat(data)
	if err != nil {
		return 0, err
	}
	return userLikeStat.LikeCount, nil
}
