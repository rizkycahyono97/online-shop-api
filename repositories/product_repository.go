package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type ProductRepository interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	FindByID(id uint) (*domain.Product, error)
	FindAll() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id uint) error
}
