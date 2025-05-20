package repositories

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (c CategoryRepositoryImpl) CreateCategory(category *domain.Category) (*domain.Category, error) {
	if err := c.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
