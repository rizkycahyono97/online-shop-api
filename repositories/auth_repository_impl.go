package repositories

import (
	"errors"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db}
}

func (repo AuthRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo AuthRepositoryImpl) CreateUser(user *domain.User) error {
	return repo.db.Create(user).Error
}
