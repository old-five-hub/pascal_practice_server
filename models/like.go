package models

import "time"

const (
	LikeQuestion int = 0
	LikeComment      = 1
)

type UserLike struct {
	ID         int       `gorm:"primary_key" json:"id""`
	NickName   Account   `json:"nickName" gorm:"references:Nickname"`
	AccountID  int       `json:"accountId"`
	Question   Question  `json:"question" gorm:"references:Name"`
	QuestionID int       `json:"questionId"`
	Comment    Comment   `json:"comment" gorm:"references:ID"`
	CommentID  int       `json:"commentId"`
	LikeType   int       `json:"likeType"`
	LikeStatus int       `default:"0"json:"listStatus"`
	CreateAt   time.Time `gorm:"autoCreateTime"json:"createAt"`
	UpdateAt   time.Time `gorm:"autoUpdateTime"json:"updateAt"`
	Deleted    int       `default:"0"json:"deleted"`
}

type UserLikeStat struct {
	ID         int       `gorm:"primary_key" json:"id""`
	Question   Question  `json:"question" gorm:"references:Name"`
	QuestionID int       `json:"questionId"`
	Comment    Comment   `json:"comment" gorm:"references:ID"`
	CommentID  int       `json:"commentId"`
	LikeCount  int       `json:"likeCount"`
	CreateAt   time.Time `gorm:"autoCreateTime"json:"createAt"`
	UpdateAt   time.Time `gorm:"autoUpdateTime"json:"updateAt"`
	Deleted    int       `default:"0"json:"deleted"`
}

func LikeContent() {
	
}
