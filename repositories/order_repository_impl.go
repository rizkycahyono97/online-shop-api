package repositories

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
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
	err := r.db.Preload("OrderItems.Product").Preload("Payment").
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

func (r OrderRepositoryImpl) GetOrderByID(orderID, userID uint) (*domain.Order, error) {
	var order domain.Order

	err := r.db.
		Preload("User").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("Payment").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error

	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r OrderRepositoryImpl) UpdateOrder(order *domain.Order) (*domain.Order, error) {
	if err := r.db.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r OrderRepositoryImpl) UpdateOrderStatus(orderID uint, status string) error {
	return r.db.Model(&domain.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
