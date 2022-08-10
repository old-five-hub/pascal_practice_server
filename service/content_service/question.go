package content_service

type CreateQuestionForm struct {
	Name   string `form:"name" valid:"required;MaxSize(100)"`
	TagIds []int  `form:"tagIds"`
}
