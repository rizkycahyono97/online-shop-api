package services

import (
	"context"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
)

type UserService interface {
	Register(req *web.UserRequest) error
	Login(req *web.UserRequest) (*domain.User, error)
	GetProfile(ctx context.Context, userId uint) (*domain.User, error)
	UpdateProfile(ctx context.Context, userId uint, req *web.UserRequest) error
	LogOut(ctx context.Context, userId uint) error
}
