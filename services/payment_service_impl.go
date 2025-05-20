package services

import (
	"errors"
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"github.com/rizkycahyono97/online-shop-api/repositories"
	"time"
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

func (p PaymentServiceImpl) CreatePayment(userID uint, payment *domain.Payment) (*domain.Payment, error) {
	// validsi order milik user
	order, err := p.orderRepo.GetOrderByID(payment.OrderID, userID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	// cek status order
	switch order.Status {
	case "cancelled":
		return nil, errors.New("order is cancelled")
	case "shipped":
		if payment.Method != "cod" {
			return nil, errors.New("payment method is invalid")
		}
	case "pending":
	default:
		return nil, errors.New("order status is invalid")
	}

	// Cek apakah sudah ada payment sebelumnya
	existingPayment, _ := p.paymentRepo.GetByOrderID(payment.OrderID)
	if existingPayment != nil {
		return nil, errors.New("payment for this order already exists")
	}

	// Validasi berdasarkan metode pembayaran
	switch payment.Method {
	case "va":
		if payment.Amount != order.TotalAmount {
			return nil, errors.New("amount must match total order amount for virtual account payments")
		}
	case "cod":
		// Tidak perlu ada amount dari request, langsung diset ke 0
		payment.Amount = 0
	case "credit", "paylater":
		minAmount := order.TotalAmount / 3
		if payment.Amount < minAmount {
			return nil, errors.New("amount must be greater than or equal to minimum amount of payment")
		}
	default:
		return nil, errors.New("payment method is invalid")
	}

	// pastikan amount sesuai order
	payment.Amount = order.TotalAmount
	payment.UserID = userID
	payment.IsPaid = true
	payment.Status = "paid"

	createPayment, err := p.paymentRepo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	//Update order status jika pembayaran berhasil
	order.Status = "paid"
	order.UpdatedAt = time.Now()
	_, err = p.orderRepo.UpdateOrder(order)
	if err != nil {
		return nil, errors.New("payment created, but failed to update order status")
	}
	
	return createPayment, nil
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
