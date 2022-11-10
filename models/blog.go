package models

import (
	"time"
)

//type Article model
type Article struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"id"`
	Nickname     string    `gorm:"size:255;not null;unique" json:"nickname"`
	Title        string    `gorm:"size:200" json:"title"`
	Content      string    `gorm:"size:8000" json:"body" `
	Comment      Comment   `json:"comment"`
	CreationDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdate"`
}

//type Comment model
type Comment struct {
	Nickname     string    `gorm:"size:255;not null;unique" json:"nickname"`
	Content      string    `gorm:"size:8000" json:"body" `
	CreationDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdate"`
}

//to set table name for article
func (article *Article) TableName() string {
	return "article"
}

//to set table name for comment
func (comment *Comment) TableName() string {
	return "comment"
}

//response map for article
func (article *Article) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = article.ID
	resp["nickname"] = article.Nickname
	resp["title"] = article.Title
	resp["content"] = article.Content
	resp["createdate"] = article.CreationDate

	return resp
}

//response map for comment
func (comment *Comment) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["nickname"] = comment.Nickname
	resp["content"] = comment.Content
	resp["createdate"] = comment.CreationDate
	return resp
}
