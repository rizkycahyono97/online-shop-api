package web

type CategoryCreateRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
}
