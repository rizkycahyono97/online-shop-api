package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"net/http"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

// Create Product
func (p *ProductController) CreateProduct(c *gin.Context) {
	var req web.ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest)
	}
}
