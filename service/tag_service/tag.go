package tag_service

import "pascal_practice_server/models"

func GetAll() ([]*models.Tag, error) {
	tags, err := models.GetTags()

	if err != nil {
		return nil, err
	}

	return tags, nil
}
