package web

type AddItemToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required,gt=0"`
}

type RemoveItemFromCartRequest struct {
	Quantity uint `json:"quantity"`
}
