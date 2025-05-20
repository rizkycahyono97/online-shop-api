package web

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type PaymentUpdateRequest struct {
	Status string `json:"status"`
}

type PaymentCreateRequest struct {
	OrderID uint    `json:"order_id" binding:"required"`
	Method  string  `json:"method" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}

func PaymentRequestToModel(req PaymentCreateRequest) *domain.Payment {
	return &domain.Payment{
		OrderID: req.OrderID,
		Method:  req.Method,
		Amount:  req.Amount,
		Status:  "pending", // default status
	}
}
