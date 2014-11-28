package models

import (
	"time"
)

type Article struct {
	Id        int64
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type ArticleResp struct {
	Article *Article `json:"article"`
}

type ArticlesResp []ArticleResp

type Articles []Article

func (a Article) TableName() string {
	return "article"
}

func AllArticles() ArticlesResp {
	var arts Articles
	artsr := ArticlesResp{}
	DB.Find(&arts)
	for _, art := range arts {
		artr := ArticleResp{
			Article: &art,
		}
		artsr = append(artsr, artr)
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
