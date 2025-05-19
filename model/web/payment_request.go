package web

type PaymentUpdateRequest struct {
	Status string `json:"status"`
}

type PaymentCreateRequest struct {
	OrderID uint    `json:"order_id" binding:"required"`
	Method  string  `json:"method" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}
