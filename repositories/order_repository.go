package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type OrderRepository interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
	CreateOrderItems(items []*domain.OrderItem) error
	GetOrderByUserID(userID uint) ([]*domain.Order, error)
	GetAllOrders() ([]*domain.Order, error)
	GetOrderByID(orderID, userID uint) (*domain.Order, error)
	UpdateOrder(order *domain.Order) (*domain.Order, error)
	UpdateOrderStatus(orderID uint, status string) error
}
