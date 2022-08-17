package content_service

import "pascal_practice_server/models"

func CreateQuestion(name string, tags []models.Tag) error {
	return models.CreateQuestion(name, tags)
}

func GetQuestionList(tagIds []int, page, limit int) (models.QuestionListResult, error) {
	return models.GetQuestionList(tagIds, page, limit)
}

func GetQuestionInfo(id int) (models.Question, error) {
	return models.GetQuestionInfo(id)
}
