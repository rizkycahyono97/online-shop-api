package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/repositories"
)

type ProductServiceImpl struct {
	repo     repositories.ProductRepository
	validate *validator.Validate
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s ProductServiceImpl) CreateProduct(req *web.ProductRequest) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s ProductServiceImpl) GetProductByID(id uint) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s ProductServiceImpl) GetAllProducts() ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s ProductServiceImpl) UpdateProduct(id uint, req *web.ProductRequest) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s ProductServiceImpl) DeleteProduct(id uint) error {
	//TODO implement me
	panic("implement me")
}
