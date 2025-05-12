package repositories

import (
	"errors"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
	"time"
)

// constructor
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates new Instance of UserRepository
func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// CreateUser
func (repo *UserRepositoryImpl) CreateUser(user *domain.User) error {
	return repo.db.Create(user).Error
}

// GetUserById
func (repo *UserRepositoryImpl) GetUserById(id uint) (*domain.User, error) {
	var user domain.User
	if err := repo.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail
func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser
func (repo *UserRepositoryImpl) UpdateUser(user *domain.User) error {
	return repo.db.Save(user).Error
}

// DeleteUserById with Softdelete
func (repo *UserRepositoryImpl) DeleteUserById(user *domain.User) error {
	user.DeletedAt = time.Now()
	return repo.db.Save(user).Error
}
