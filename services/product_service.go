package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
)

type ProductService interface {
	CreateProduct(req *web.ProductRequest) (*domain.Product, error)
	GetProductByID(id uint) (*domain.Product, error)
	GetAllProducts() ([]*domain.Product, error)
	UpdateProduct(id uint, req *web.ProductRequest) (*domain.Product, error)
	DeleteProduct(id uint) error
}
