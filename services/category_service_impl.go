package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/repositories"
	"time"
)

type CategoryServiceImpl struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{repo: repo}
}

func (c CategoryServiceImpl) CreateCategory(req *web.CategoryCreateRequest) (*domain.Category, error) {
	category := &domain.Category{
		CategoryName: req.CategoryName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return c.repo.CreateCategory(category)
}
