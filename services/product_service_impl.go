package services

import (
	"errors"
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
	// validasi input
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL,
	}

	return s.repo.CreateProduct(product)
}

func (s ProductServiceImpl) GetProductByID(id uint) (*domain.Product, error) {
	return s.repo.FindProductByID(id)
}

func (s ProductServiceImpl) GetAllProducts() ([]*domain.Product, error) {
	return s.repo.FindAllProducts()
}

func (s ProductServiceImpl) UpdateProduct(id uint, req *web.ProductRequest) (*domain.Product, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// cek produk
	existing, err := s.repo.FindProductByID(id)
	if err != nil {
		return nil, errors.New("Product not Found")
	}

	// update field
	existing.Name = req.Name
	existing.Description = req.Description
	existing.Price = req.Price
	existing.Stock = req.Stock
	existing.ImageURL = req.ImageURL

	return s.repo.UpdateProduct(existing)
}

func (s ProductServiceImpl) DeleteProduct(id uint) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return errors.New("Product not Found")
	}

	return s.repo.DeleteProduct(id)
}
