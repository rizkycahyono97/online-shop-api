package web

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"time"
)

type PaymentResponse struct {
	ID        uint      `json:"id"`
	OrderID   uint      `json:"order_id"`
	Method    string    `json:"method"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PaymentResponseFromModel(payment *domain.Payment) *PaymentResponse {
	return &PaymentResponse{
		ID:        payment.ID,
		OrderID:   payment.OrderID,
		Method:    payment.Method,
		Amount:    payment.Amount,
		Status:    payment.Status,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}
}

func PaymentResponseFromModels(payments []*domain.Payment) []*PaymentResponse {
	var response []*PaymentResponse
	for _, payment := range payments {
		response = append(response, PaymentResponseFromModel(payment))
	}
	return response
}
