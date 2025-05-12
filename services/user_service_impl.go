package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/repositories"
)

type UserServiceImpl struct {
	repo     repositories.UserRepository
	validate *validator.Validate
}

// Constructor for UserService
func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *UserServiceImpl) GetProfile(userId uint) (*domain.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *UserServiceImpl) UpdateProfile(userId uint, req *web.UserRequest) error {
	// validate input
	if err := s.validate.Var(req.Name, "required,min=3"); err != nil {
		return errors.New("Invalid Name")
	}
	if err := s.validate.Var(req.Email, "required,email"); err != nil {
		return errors.New("Invalid Email")
	}

	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return err
	}

	// Prepare updates
	user.Name = req.Name
	user.Email = req.Email

	return s.repo.UpdateUser(user)
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, userId uint) error {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return errors.New("User Not Found")
	}

	return s.repo.DeleteUserById(user)
}
