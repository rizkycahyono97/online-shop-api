package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type AuthRepository interface {
	FindByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) error
}
