package service

import (
	"JET_DEV_TASK/JetDev/api/repository"
	"JET_DEV_TASK/JetDev/models"
)

type BlogService struct {
	repository repository.BlogRepository
}

func NewBlogService(r repository.BlogRepository) BlogService {
	return BlogService{
		repository: r,
	}
}

//List all
func (b BlogService) ListAll(article models.Article) (*[]models.Article, int64, error) {
	return b.repository.ListAll(article)
}

//Get content
func (b BlogService) ListContent(article models.Article, keyword string) (*[]models.Article, error) {
	return b.repository.ListContent(article, keyword)
}

//Post article
func (b BlogService) PostArticle(article models.Article, keyword string) error {
	return nil
}

//List all comments on article
func (b BlogService) ListComment(article models.Article, keyword string) (models.Article, error) {
	return b.repository.ListComment(article, keyword)
}
