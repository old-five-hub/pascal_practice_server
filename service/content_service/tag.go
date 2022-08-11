package content_service

import "pascal_practice_server/models"

type Tag struct {
	Id   int
	Name string
	Hot  int
}

func GetAllTags() ([]models.Tag, error) {
	tags, err := models.GetAllTags()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func CreateTag(t *Tag) error {
	return models.CreateTag(t.Name, t.Hot)
}

func UpdateTag(t *Tag) error {
	return models.EditTag(t.Id, t)
}

func DeleteTag(id int) error {
	return models.DeleteTag(id)
}

func ExistTagByName(name string) (bool, error) {
	return models.ExistTagByName(name)
}

func ExistTagById(id int) (bool, error) {
	return models.ExistTagById(id)
}

func ExistTagByIds(ids []int) (bool, error) {
	return models.ExistTagByIds(ids)
}

func GetTagByIds(ids []int) ([]models.Tag, error) {
	return models.GetTagByIds(ids)
}
