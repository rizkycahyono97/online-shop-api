package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
)

type AuthService interface {
	Register(req *web.UserRequest) (*domain.User, error)
	Login(req *web.UserRequest) (string, error)
}
