package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/model/web"
)

type OrderService interface {
	CreateOrder(userID uint, req *web.CreateOrderRequest) (*domain.Order, error)
	GetOrderByID(orderID, userID uint) (*web.OrderResponse, error)
	GetOrderByUserID(userID uint) ([]*domain.Order, error)
	GetAllOrders() ([]*domain.Order, error)
	GetAllUserOrders() ([]*web.OrderResponse, error)
	CancelOrder(UserID uint, orderID uint) error
	ConfirmOrder(userID uint, orderID uint) error
	UpdateOrderStatus(orderID uint, status string) (*domain.Order, error)
}
