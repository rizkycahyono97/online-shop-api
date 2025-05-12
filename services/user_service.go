package services

import (
	"context"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
)

type UserService interface {
	Register(req *web.UserRequest) error
	Login(req *web.UserRequest) (*domain.User, error)
	GetProfile(userId uint) (*domain.User, error)
	UpdateProfile(userId uint, req *web.UserRequest) error
	DeleteUser(ctx context.Context, userId uint) error
}
