package services

import (
	"errors"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/repositories"
)

type PaymentServiceImpl struct {
	paymentRepo repositories.PaymentRepository
	orderRepo   repositories.OrderRepository
}

func NewPaymentService(paymentRepo repositories.PaymentRepository, orderRepo repositories.OrderRepository) PaymentService {
	return &PaymentServiceImpl{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
	}
}

func (p PaymentServiceImpl) CreatePayment(payment *domain.Payment) (*domain.Payment, error) {
	order, err := p.orderRepo.GetOrderByID(payment.OrderID, payment.OrderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	// pastika tidak ada payment sebelumnya
	existingPayment, _ := p.orderRepo.GetOrderByID(payment.OrderID, payment.OrderID)
	if existingPayment != nil {
		return nil, errors.New("order already exists")
	}

	// pastikan amount sesuai order
	payment.Amount = order.TotalAmount

	return p.paymentRepo.CreatePayment(payment)
}

func (p PaymentServiceImpl) UpdatePaymentStatus(orderID uint, status string) (*domain.Payment, error) {
	payment, err := p.paymentRepo.GetByOrderID(orderID)
	if err != nil {
		return nil, errors.New("payment not found")
	}

	payment.Status = status
	return p.paymentRepo.UpdatePayment(payment)
}

func (p PaymentServiceImpl) GetPaymentByOrderID(orderID uint) (*domain.Payment, error) {
	return p.paymentRepo.GetByOrderID(orderID)
}

func (p PaymentServiceImpl) GetAllPayment() ([]*domain.Payment, error) {
	return p.paymentRepo.GetAllPayments()
}

func (p PaymentServiceImpl) GetPaymentByUserID(userID uint) ([]*domain.Payment, error) {
	return p.paymentRepo.GetPaymentByUserID(userID)
}
