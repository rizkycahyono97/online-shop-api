package repositories

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (o OrderRepositoryImpl) CreateOrder(order *domain.Order) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepositoryImpl) CreateOrderItems(items []*domain.OrderItem) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepositoryImpl) GetOrderByUserID(userID uint) ([]*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepositoryImpl) GetAllOrders() ([]*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}
