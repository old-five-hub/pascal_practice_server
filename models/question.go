package models

import "time"

type Question struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Tags     []Tag     `gorm:"many2many:question_tag"json:"tags"`
	CreateAt time.Time `gorm:"autoCreateTime"json:"createAt"`
	UpdateAt time.Time `gorm:"autoUpdateTime"json:"updateAt"`
	Deleted  int       `default:"0"json:"deleted"`
}

func CreateQuestion(name string, tags []Tag) error {
	question := Question{
		Name: name,
		Tags: tags,
	}
	return db.Create(&question).Error
}
