package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

// kontrak user_repository_impl
type UserRepository interface {
	GetUsers() ([]*domain.User, error)
	GetUserById(id uint) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUserById(user *domain.User) error
}
