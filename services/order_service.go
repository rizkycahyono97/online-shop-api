package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
)

type OrderService interface {
	CreateOrder(userID uint, req *web.CreateOrderRequest) (*domain.Order, error)
	GetOrderByID(orderID uint) (*domain.Order, error)
	GetOrderByUserID(userID uint) ([]*domain.Order, error)
	GetAllOrders() ([]*domain.Order, error)
	CancelOrder(UserID uint, orderID uint) error
	ConfirmOrder(userID uint, orderID uint) error
}
