package models

import (
	"time"
)

type Article struct {
	Id        int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Articles []Article

func (a Article) TableName() string {
	return "article"
}

func AllArticles() []Article {
	var arts Articles
	DB.Find(&arts)
	return arts
}
