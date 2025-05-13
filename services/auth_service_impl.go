package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rizkycahyono97/online-shop-api/config"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/repositories"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthServiceImpl struct {
	repo repositories.AuthRepository
}

func NewAuthServiceImpl(repo repositories.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s AuthServiceImpl) Register(req *web.UserRequest) (*domain.User, error) {
	existing, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("User already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "customer",
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s AuthServiceImpl) Login(req *web.UserRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("Email not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("Password Incorect")
	}

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	secret := config.GetEnv("JWT_SECRET", "")
	if secret == "" {
		return "", errors.New("JWT_SECRET environment variable not set")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("Error signing token")
	}

	return tokenString, nil
}
