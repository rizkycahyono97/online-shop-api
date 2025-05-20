package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
)

type CategoryService interface {
	CreateCategory(req *web.CategoryCreateRequest) (*domain.Category, error)
}
