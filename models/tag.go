package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Hot      int       `json:"hot"`
	CreateAt time.Time `gorm:"autoCreateTime"json:"createAt"`
	UpdateAt time.Time `gorm:"autoUpdateTime"json:"updateAt"`
	Deleted  int       `default:"0"json:"deleted"`
}

func GetAllTags() ([]Tag, error) {
	var tags []Tag
	result := db.Find(&tags)

	if result.Error != nil {
		return nil, result.Error
	}
	return tags, nil
}

func CreateTag(name string, hot int) error {
	tag := Tag{
		Name: name,
		Hot:  hot,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Update("delete", 1).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? And deleted = 0", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.Id > 0 {
		return true, nil
	}
	return false, nil
}

func ExistTagById(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? And deleted = 0", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.Id > 0 {
		return true, nil
	}
	return false, nil
}

func ExistTagByIds(ids []int) (bool, error) {
	var tags []Tag
	err := db.Select("id").Where("id in (?) And deleted = 0", ids).Find(&tags).Error
	if err != nil {
		return false, err
	}
	if len(tags) == len(ids) {
		return true, nil
	}
	return false, nil
}

func GetTagByIds(ids []int) ([]Tag, error) {
	var tags []Tag
	err := db.Select("id").Where("id in (?) And deleted = 0", ids).Find(&tags).Error

	if err != nil {
		return nil, err
	}
	return tags, nil
}
