package services

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type ProductService interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	GetProductByID(id uint) (*domain.Product, error)
	GetAllProducts() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id uint) error
}
