package models

type Book struct {
	Id           int    `json:"id"`
	Title        string `json:"title" validate:"required"`
	Descr        string `json:"description"`
	ThumbnailUrl string `json:"thumbnail_url" validate:"required"`
}
