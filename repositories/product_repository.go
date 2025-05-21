package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type ProductRepository interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	FindProductByID(id uint) (*domain.Product, error)
	FindAllProducts() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id uint) error
	SearchProductByKeyword(keyword string) ([]*domain.Product, error)
}
