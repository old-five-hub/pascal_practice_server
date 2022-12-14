package models

import (
	"time"
)

type Question struct {
	Id       int       `gorm:"primary_key"json:"id"`
	Name     string    `json:"name"`
	Tags     []Tag     `gorm:"many2many:question_tag"json:"tags"`
	Comments []Comment `json:"comments"`
	Desc     string    `json:"desc"`
	Answer   string    `json:"answer"`
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

type QuestionListResult struct {
	List    []Question
	Total   int64
	HasMore bool
}

func GetQuestionList(tagIds []int, page, limit int) (QuestionListResult, error) {
	var questions []Question
	var total int64
	var hasMore bool
	var err error

	if len(tagIds) > 0 {
		err = db.Model(&Question{}).Where("id IN (SELECT question_id FROM question_tag where tag_id in (?))", tagIds).Count(&total).Error

		if err != nil {
			return QuestionListResult{}, err
		}

		err = db.Preload("Tags").
			Where("id IN (SELECT question_id FROM question_tag where tag_id in (?))", tagIds).
			Offset((page - 1) * limit).
			Limit(limit).
			Find(&questions).Error
	} else {
		err = db.Model(&Question{}).Count(&total).Error

		if err != nil {
			return QuestionListResult{}, err
		}

		err = db.Preload("Tags").
			Offset((page - 1) * limit).
			Limit(limit).
			Find(&questions).Error
	}

	if err != nil {
		return QuestionListResult{}, err
	}

	hasMore = limit*page < int(total)

	return QuestionListResult{
		List:    questions,
		Total:   total,
		HasMore: hasMore,
	}, nil
}

func GetQuestionInfo(id int) (Question, error) {
	question := Question{
		Id: id,
	}
	err := db.Preload("Tags").First(&question).Error
	if err != nil {
		return question, err
	}

	return question, nil
}
