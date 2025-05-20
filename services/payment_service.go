package services

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type PaymentService interface {
	CreatePayment(userID uint, payment *domain.Payment) (*domain.Payment, error)
	UpdatePaymentStatus(orderID uint, status string) (*domain.Payment, error)
	GetPaymentByOrderID(orderID uint) (*domain.Payment, error)
	GetAllPayment() ([]*domain.Payment, error)
	GetPaymentsForUser(userID uint) ([]*domain.Payment, error)
}
