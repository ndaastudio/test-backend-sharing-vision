package service

import (
	"errors"
	"test-backend-sharing-vision/model"
	"test-backend-sharing-vision/repository"
	"test-backend-sharing-vision/request"
	"time"
)

type ArticleService interface {
	GetAll(limit, offset int, status string) ([]model.Post, error)
	GetByID(id uint) (model.Post, error)
	Create(req request.CreateArticleRequest) (model.Post, error)
	Update(id uint, req request.UpdateArticleRequest) (model.Post, error)
	Delete(id uint) error
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) GetAll(limit, offset int, status string) ([]model.Post, error) {
	return s.repo.FindAll(limit, offset, status)
}

func (s *articleService) GetByID(id uint) (model.Post, error) {
	return s.repo.FindByID(id)
}

func (s *articleService) Create(req request.CreateArticleRequest) (model.Post, error) {
	article := model.Post{
		Title:       req.Title,
		Content:     req.Content,
		Category:    req.Category,
		Status:      req.Status,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	return s.repo.Create(article)
}

func (s *articleService) Update(id uint, req request.UpdateArticleRequest) (model.Post, error) {
	article, err := s.repo.FindByID(id)
	if err != nil {
		return article, errors.New("Artikel tidak ditemukan")
	}

	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Content != "" {
		article.Content = req.Content
	}
	if req.Category != "" {
		article.Category = req.Category
	}
	if req.Status != "" {
		article.Status = req.Status
	}
	article.UpdatedDate = time.Now()

	return s.repo.Update(article)
}

func (s *articleService) Delete(id uint) error {
	article, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("Artikel tidak ditemukan")
	}
	return s.repo.Delete(article)
}
