package repositories

import (
	"errors"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (repo ProductRepositoryImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	if err := repo.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo ProductRepositoryImpl) FindProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	if err := repo.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &product, nil
}

func (repo ProductRepositoryImpl) FindAllProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	if err := repo.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}

func (repo ProductRepositoryImpl) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	if err := repo.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo ProductRepositoryImpl) DeleteProduct(id uint) error {
	if err := repo.db.Delete(&domain.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
