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

type CartAdminResponse struct {
	CartID   uint               `json:"cart_id"`
	UserID   uint               `json:"user_id"`
	UserName string             `json:"user_name"`
	Email    string             `json:"email"`
	Items    []CartItemResponse `json:"items"`
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

// function helper untuk mengembalikan semua carts, users, dan items
func CartAdminResponseFromModel(cart domain.Cart) CartAdminResponse {
	var items []CartItemResponse
	for _, item := range cart.CartItems {
		items = append(items, CartItemResponse{
			ProductID: item.ProductID,
			Name:      item.Product.Name,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
			ImageURL:  item.Product.ImageURL,
		})
	}

	return CartAdminResponse{
		CartID:   cart.ID,
		UserID:   uint(cart.User.ID),
		UserName: cart.User.Name,
		Email:    cart.User.Email,
		Items:    items,
	}
}

// function helper untuk slice mengembalikan semua carts, users, dan items
func CartAdminResponseFromModels(carts []*domain.Cart) []CartAdminResponse {
	var responses []CartAdminResponse
	for _, cart := range carts {
		responses = append(responses, CartAdminResponseFromModel(*cart))
	}
	return responses
}
