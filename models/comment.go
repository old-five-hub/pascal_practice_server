package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID         int       `gorm:"primary_key" json:"id""`
	TopicType  int       `default:"0"json:"topicType"`
	Content    string    `json:"content"`
	Question   Question  `json:"question"`
	QuestionID int       `gorm:"foreignKey"json:"questionId"`
	Author     Account   `json:"author"gorm:"foreignKey:AuthorID"`
	AuthorID   int       `json:"AuthorID"`
	ParentID   int       `json:"parentId"`
	CreateAt   time.Time `gorm:"autoCreateTime"json:"createAt"`
	UpdateAt   time.Time `gorm:"autoUpdateTime"json:"updateAt"`
	Deleted    int       `default:"0"json:"deleted"`
}

func GetComments(questionId int) ([]Comment, error) {
	var comments []Comment
	err := db.Preload("Author", func(query *gorm.DB) *gorm.DB {
		return query.Select("id, nickname, avatar, follow")
	}).Find(&comments).Where("question_id = ?", questionId).Error
	if err != nil {
		return comments, nil
	}
	return comments, nil
}

func CreateComment(data map[string]interface{}) error {
	question := Question{
		Id: data["questionId"].(int),
	}
	comment := Comment{
		QuestionID: data["questionId"].(int),
		Content:    data["content"].(string),
		ParentID:   data["parentId"].(int),
	}
	err := db.Model(&question).Association("Comments").Append(comment)

	if err != nil {
		return err
	}
	return nil
}
