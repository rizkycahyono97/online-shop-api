package web

type CreateOrderRequest struct {
	Address string `json:"address" binding:"required"`
}
