package content_service

import "pascal_practice_server/models"

func GetComments(questionId int) ([]models.Comment, error) {
	return models.GetComments(questionId)
}

func CreateComment(data map[string]interface{}) error {
	return models.CreateComment(data)
}
