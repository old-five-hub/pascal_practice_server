package models

import (
	"log"
)

type Tag struct {
	Id       int
	Name     string
	Hot      int
	CreateAt string
	UpdateAt string
}

func GetTags() ([]*Tag, error) {
	var tags []*Tag
	rows, err := Db.Query("select * from tag")

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(Tag)
		err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.CreateAt, &tmp.UpdateAt, &tmp.Hot)
		if err != nil {
			log.Fatal(err)
		}
		tags = append(tags, tmp)
		log.Println(tags)
	}

	return tags, err
}
