package content_service

import "pascal_practice_server/models"

func UpdateUserLike(id int, data map[string]interface{}) error {
	return models.UpdateUserLike(id, data)
}

func ExistUserLike(data map[string]interface{}) (models.UserLike, error) {
	return models.ExistUserLike(data)
}

func CreateUserLike(data map[string]interface{}) error {
	return models.CreateUserLike(data)
}

func UpdateUserLikeStat(id int, data map[string]interface{}) error {
	return models.UpdateUserLikeStat(id, data)
}

func ExistUserLikeStat(data map[string]interface{}) (models.UserLikeStat, error) {
	return models.ExistUserLikeStat(data)
}

func CreateUserLikeStat(data map[string]interface{}) error {
	return models.CreateUserLikeStat(data)
}
