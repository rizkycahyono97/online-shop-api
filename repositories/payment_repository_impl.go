package repositories

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}

func (p PaymentRepositoryImpl) CreatePayment(payment *domain.Payment) (*domain.Payment, error) {
	if err := p.db.Debug().Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (p PaymentRepositoryImpl) UpdatePayment(payment *domain.Payment) (*domain.Payment, error) {
	if err := p.db.Save(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (p PaymentRepositoryImpl) GetByOrderID(orderID uint) (*domain.Payment, error) {
	var payment domain.Payment
	err := p.db.Where("order_id = ?", orderID).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (p PaymentRepositoryImpl) GetAllPayments() ([]*domain.Payment, error) {
	var payments []*domain.Payment
	err := p.db.Order("created_at desc").Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (p PaymentRepositoryImpl) GetPaymentByUserID(userID uint) ([]*domain.Payment, error) {
	var payments []*domain.Payment
	err := p.db.Joins("JOIN orders ON orders.id = payments.order_id").
		Where("orders.user_id = ?", userID).
		Order("payments.created_at desc").
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}
