package services

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type OrderService interface {
	CreateOrder(userID uint) (*domain.Order, error)
	GetOrderByID(orderID uint) (*domain.Order, error)
	GetOrderByUserID(userID uint) ([]*domain.Order, error)
	GetAllOrders() ([]*domain.Order, error)
	CancelOrder(UserID uint) (*domain.Order, error)
}
