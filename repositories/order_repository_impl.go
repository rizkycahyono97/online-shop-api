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

func (r OrderRepositoryImpl) CreateOrder(order *domain.Order) (*domain.Order, error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r OrderRepositoryImpl) CreateOrderItems(items []*domain.OrderItem) error {
	if err := r.db.Create(items).Error; err != nil {
		return err
	}
	return nil
}

func (r OrderRepositoryImpl) GetOrderByUserID(userID uint) ([]*domain.Order, error) {
	var orders []*domain.Order
	err := r.db.Preload("OrderItems.Product").Preload("payment").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r OrderRepositoryImpl) GetAllOrders() ([]*domain.Order, error) {
	var orders []*domain.Order
	err := r.db.Preload("User").Preload("OrderItems.Product").Preload("payment").
		Order("created_at desc").
		Find(&orders).Error

	if err != nil {
		return nil, err
	}
	return orders, nil
}
