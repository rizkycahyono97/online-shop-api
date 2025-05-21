package web

import (
	"github.com/rizkycahyono97/online-shop-api/model/domain"
	"time"
)

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       uint      `json:"stock"`
	ImageURL    string    `json:"image_url"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ProductResponseFromModel(product *domain.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ProductResponseListFromModel(products []*domain.Product) []ProductResponse {
	var response []ProductResponse
	for _, p := range products {
		response = append(response, ProductResponseFromModel(p))
	}
	return response
}
