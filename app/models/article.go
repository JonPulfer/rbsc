package models

import (
	"github.com/lib/pq"
	"time"
)

type Article struct {
	Id        int64       `json:"id"`
	Title     string      `json:"title"`
	Body      string      `json:"body"`
	CreatedAt pq.NullTime `json:"-"`
	UpdatedAt pq.NullTime `json:"-"`
	DeletedAt pq.NullTime `json:"-"`
}

type ArticleResp struct {
	Article *Article `json:"article"`
}

type ArticlesResp struct {
	Articles Articles `json:"articles"`
}

type Articles []Article

func (a Article) TableName() string {
	return "article"
}

func (art *Article) BeforeCreate() (err error) {
	art.CreatedAt.Time = time.Now()
	art.CreatedAt.Valid = true
	art.UpdatedAt.Time = time.Now()
	art.UpdatedAt.Valid = true
	return
}
func (art *Article) BeforeSave() (err error) {
	art.UpdatedAt.Time = time.Now()
	art.UpdatedAt.Valid = true
	return
}

func AllArticles() ArticlesResp {
	var arts Articles
	DB.Find(&arts)
	artsr := ArticlesResp{
		Articles: arts,
	}
	return artsr
}

func ArticleById(id int) ArticleResp {
	var art Article
	DB.First(&art, id)
	artr := ArticleResp{
		Article: &art,
	}
	return artr
}
