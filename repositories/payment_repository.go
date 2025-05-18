package repositories

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type PaymentRepository interface {
	CreatePayment(payment *domain.Payment) (*domain.Payment, error)
	UpdatePayment(payment *domain.Payment) (*domain.Payment, error)
	GetByOrderID(orderID uint) (*domain.Payment, error)
	GetAllPayments() ([]*domain.Payment, error)
	GetPaymentByUserID(userID uint) ([]*domain.Payment, error)
}
