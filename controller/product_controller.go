package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"strconv"
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
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	product, err := p.productService.CreateProduct(&req)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Product Created Successfully", web.ProductResponseFromModel(product))
}

// Get All Product
func (p *ProductController) GetAllProducts(c *gin.Context) {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Products Fetch Successfully", products)
}

// Get Product By ID
func (p *ProductController) GetProductById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid Product ID", nil)
		return
	}

	product, err := p.productService.GetProductByID(uint(id))
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "ID Not Found", nil)
		return
	}

	helpers.JSONSuccessResponse(c, "Product Fetch Successfully", web.ProductResponseFromModel(product))
}

// Update
func (p *ProductController) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid Product ID", nil)
		return
	}

	var req *web.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	product, err := p.productService.UpdateProduct(uint(id), req)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Product Updated Successfully", web.ProductResponseFromModel(product))
}

// Delete
func (p *ProductController) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid Product ID", nil)
		return
	}

	err = p.productService.DeleteProduct(uint(id))
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Product Deleted Successfully", nil)
}

// SearchProductByKeyword
func (p *ProductController) SearchProducts(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		helpers.JSONBadRequestResponse(c, "Invalid Search Keyword", nil)
		return
	}

	products, err := p.productService.SearchProduct(keyword)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Products Search Successfully", web.ProductResponseListFromModel(products))
}
