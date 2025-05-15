package web

import "github.com/rizkycahyono97/online-shop-api/model/domain"

type CartItemResponse struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	ImageURL  string  `json:"image_url"`
}

// function helper untuk mengembalikan cartitemsRequest dengan slice
func CartItemsResponseFromModels(items []*domain.CartItems) []CartItemResponse {
	var responses []CartItemResponse
	for _, item := range items {
		response := CartItemResponse{
			ID:        item.Product.ID,
			ProductID: item.ProductID,
			Name:      item.Product.Name,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
			ImageURL:  item.Product.ImageURL,
		}
		responses = append(responses, response)
	}
	return responses
}

// function helper untuk mengembalikan cartitemsRequest untuk 1 item saja
func CartItemResponseFromModels(item *domain.CartItems) CartItemResponse {
	return CartItemResponse{
		ID:        item.ID,
		ProductID: item.ProductID,
		Name:      item.Product.Name,
		Quantity:  item.Quantity,
		Price:     item.Product.Price,
		ImageURL:  item.Product.ImageURL,
	}
}
