package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"net/http"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var req web.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
		})
		return
	}

	category, err := cc.service.CreateCategory(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Category created",
		Data:    category,
	})
}
