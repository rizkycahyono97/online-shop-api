package web

type CreateOrderRequest struct {
	Address string `json:"address" binding:"required"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending paid shipped cancelled"`
}
