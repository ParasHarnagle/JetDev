package repository

import (
	"JET_DEV_TASK/JetDev/infrastructure"
	"JET_DEV_TASK/JetDev/models"
)

//Blog repository
type BlogRepository struct {
	db infrastructure.Database
}

func NewBlogRepository(db infrastructure.Database) BlogRepository {
	return BlogRepository{
		db: db,
	}
}

type CommentRepository struct {
	db infrastructure.Database
}

func NewCommentRepository(db infrastructure.Database) CommentRepository {
	return CommentRepository{
		db: db,
	}
}

//Save repo for article
func (b BlogRepository) PostArticle(blog models.Article) error {
	return b.db.DB.Create(&blog).Error
}

//Save repo for comment
func (b BlogRepository) SaveComment(blog models.Comment) error {
	return b.db.DB.Create(&blog).Error
}

//List all articles
func (b BlogRepository) ListAll(article models.Article) (*[]models.Article, int64, error) {
	var articles []models.Article
	var totalRows int64 = 0
	// Get all records
	result := b.db.DB.Limit(20).Find(&article)
	// SELECT * FROM articles;
	err := result.
		Where(article).
		Find(&articles).
		Count(&totalRows).Error
	return &articles, totalRows, err

}

//get an article content
func (b BlogRepository) ListContent(article models.Article, keyword string) (*[]models.Article, error) {
	var articles []models.Article

	result := b.db.DB.Select("article.Content").Where("article.Title LIKE ?", keyword)
	err := result.
		Debug().
		Model(&models.Article{}).
		Where(&article).
		Take(&articles).Error
	return &articles, err
}

//List all comments on article by nickname
func (b BlogRepository) ListComment(article models.Article, keyword string) (models.Article, error) {
	var articles models.Article

	result := b.db.DB.Select("article.Comment").Where("article.Nickname LIKE ?", keyword)
	err := result.
		Debug().
		Model(&models.Article{}).
		Where(&article).
		Take(&articles).Error
	return articles, err
}
