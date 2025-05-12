package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/repositories"
	"golang.org/x/crypto/bcrypt"
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

// Register new User
func (s *UserServiceImpl) Register(req *web.UserRequest) error {
	// validate input
	if err := s.validate.Var(req.Name, "required,min=3"); err != nil {
		return errors.New("Invalid Name")
	}
	if err := s.validate.Var(req.Email, "required,email"); err != nil {
		return errors.New("Invalid Email")
	}
	if err := s.validate.Var(req.Password, "required,min="); err != nil {
		return errors.New("Invalid Password")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Failed to Generate Password")
	}

	// Create user model
	newUser := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "customer", // default role
	}

	// simpan ke database
	err = s.repo.CreateUser(newUser)
	if err != nil {
		return err
	}

	// Save user to database
	return s.repo.CreateUser(newUser)
}
