package web

type ProductRequest struct {
	Name        string  `json:"name" binding:"required,max=255" validate:"required,min=3"`
	Description string  `json:"description" binding:"required" validate:"required"`
	Price       float64 `json:"price" binding:"required" validate:"required,gt=0"`
	Stock       uint    `json:"stock" binding:"required" validate:"required,gte=0"`
	ImageURL    string  `json:"image_url" binding:"required" validate:"required,url"`
}
