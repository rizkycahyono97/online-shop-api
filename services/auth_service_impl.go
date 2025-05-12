package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/repositories"
)

type AuthServiceImpl struct {
	repo repositories.AuthRepository
}

func NewAuthServiceImpl(repo repositories.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (a AuthServiceImpl) Register(req *web.UserRequest) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) Login(req *web.UserRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

//Register User
