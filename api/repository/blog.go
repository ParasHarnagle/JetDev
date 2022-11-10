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
func (b BlogRepository) SaveArticle(blog models.Article) error {
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

//get an article content based on title
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

//List all comments on article
func (b BlogRepository) ListComment(article models.Article, keyword string) (*[]models.Article, error) {
	var articles models.Article

	result := b.db.DB.Select("article.Comment").Find(&articles)
	err := result.
		Debug().
		Model(&models.Article{}).
		Where(&article).
		Take(&articles).Error
	return articles, err
}

//POst an article
func (b BlogRepository) PostArticle(article models.Article, keyword string) (*[]models.Article, error) {
	return b.repository.PostArticle(article, keyword)
}

//comment on comment
