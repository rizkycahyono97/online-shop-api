package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
)

type UserService interface {
	GetProfile(userId uint) (*domain.User, error)
	GetAllProfiles() ([]*domain.User, error)
	UpdateProfile(userId uint, req *web.UserRequest) (*domain.User, error)
	DeleteProfile(userId uint) error
}
