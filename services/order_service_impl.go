package services

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/repositories"
)

type OrderServiceImpl struct {
	orderRepo   repositories.OrderRepository
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
}

func NewOrderService(
	orderRepo repositories.OrderRepository,
	cartRepo repositories.CartRepository,
	productRepo repositories.ProductRepository,
) OrderService {
	return &OrderServiceImpl{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (o OrderServiceImpl) CreateOrder(userID uint) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderServiceImpl) GetOrderByID(orderID uint) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderServiceImpl) GetOrderByUserID(userID uint) ([]*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderServiceImpl) GetAllOrders() ([]*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderServiceImpl) CancelOrder(UserID uint) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}
