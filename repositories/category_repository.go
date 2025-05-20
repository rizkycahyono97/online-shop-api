package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type CategoryRepository interface {
	CreateCategory(category *domain.Category) (*domain.Category, error)
}
