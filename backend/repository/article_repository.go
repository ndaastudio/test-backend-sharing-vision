package repository

import (
	"test-backend-sharing-vision/model"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll(limit, offset int, status string) ([]model.Post, error)
	FindByID(id uint) (model.Post, error)
	Create(article model.Post) (model.Post, error)
	Update(article model.Post) (model.Post, error)
	Delete(article model.Post) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) FindAll(limit, offset int, status string) ([]model.Post, error) {
	var articles []model.Post
	err := r.db.Limit(limit).Offset(offset).Where("status = ?", status).Find(&articles).Error
	return articles, err
}

func (r *articleRepository) FindByID(id uint) (model.Post, error) {
	var article model.Post
	err := r.db.First(&article, id).Error
	return article, err
}

func (r *articleRepository) Create(article model.Post) (model.Post, error) {
	err := r.db.Create(&article).Error
	return article, err
}

func (r *articleRepository) Update(article model.Post) (model.Post, error) {
	err := r.db.Save(&article).Error
	return article, err
}

func (r *articleRepository) Delete(article model.Post) error {
	return r.db.Delete(&article).Error
}
