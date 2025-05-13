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
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// GetAllUsers
func (repo *UserRepositoryImpl) GetUsers() ([]*domain.User, error) {
	var users []*domain.User

	err := repo.db.Where("deleted_at = ?", "0000-00-00 00:00:00").Find(&users).Error
	return users, err
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

// UpdateUser
func (repo *UserRepositoryImpl) UpdateUser(user *domain.User) error {
	return repo.db.Save(user).Error
}

// DeleteUserById with Softdelete
func (repo *UserRepositoryImpl) DeleteUserById(user *domain.User) error {
	user.DeletedAt = time.Now()
	return repo.db.Save(user).Error
}
