package services

import (
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

// Get Profile by id
func (s *UserServiceImpl) GetProfile(userId uint) (*domain.User, error) {
	return s.repo.GetUserById(userId)
}

// Get All Profile
func (s *UserServiceImpl) GetAllProfiles() ([]*domain.User, error) {
	return s.repo.GetUsers()
}

// Update Profile
func (s *UserServiceImpl) UpdateProfile(userId uint, req *web.UserRequest) (*domain.User, error) {
	// Validasi input
	if err := s.validate.Var(req.Name, "required,min=3"); err != nil {
		return nil, errors.New("Invalid Name")
	}
	if err := s.validate.Var(req.Email, "required,email"); err != nil {
		return nil, errors.New("Invalid Email")
	}

	// Cari user
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// Update data
	user.Name = req.Name
	user.Email = req.Email

	err = s.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete User Profile
func (s *UserServiceImpl) DeleteProfile(userId uint) error {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return errors.New("User Not Found")
	}

	return s.repo.DeleteUserById(user)
}
