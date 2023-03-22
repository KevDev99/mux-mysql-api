package models

import "reflect"

type Book struct {
	Id           int    `json:"id" gorm:"column:book_id"`
	Title        string `json:"title" validate:"required"`
	Descr        string `json:"descr"`
	ThumbnailUrl string `json:"thumbnail_url" validate:"required"`
	CreatedAt    string `json:"created_at" validate:"required"`
}

func (b Book) GetJsonFields() []string {
	var jsonFields []string
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		jsonFields = append(jsonFields, val.Type().Field(i).Tag.Get("json"))

	}

	return jsonFields
}
