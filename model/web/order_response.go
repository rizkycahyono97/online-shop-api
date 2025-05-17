package web

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"time"
)

// Response Order Item
type OrderItemResponse struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}

// Response Order Response
type OrderResponse struct {
	ID          uint                `json:"id"`
	UserID      uint                `json:"user_id"`
	TotalAmount float64             `json:"total_amount"`
	Status      string              `json:"status"`
	Address     string              `json:"address"`
	CreatedAt   time.Time           `json:"created_at"`
	Items       []OrderItemResponse `json:"items"`
}

// helper for Order Response
func OrderResponseFromModels(order *domain.Order) *OrderResponse {
	items := make([]OrderItemResponse, 0)
	for _, item := range order.OrderItems {
		items = append(items, OrderItemResponse{
			ProductID: item.ProductID,
			Name:      item.Product.Name,
			Quantity:  item.Quantity,
			Price:     float64(item.Price),
		})
	}

	return &OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		Address:     order.Address,
		CreatedAt:   order.CreatedAt,
		Items:       items,
	}
}
