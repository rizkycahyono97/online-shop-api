package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
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
		helpers.JSONSuccessResponse(c, "", err)
		return
	}

	category, err := cc.service.CreateCategory(&req)
	if err != nil {
		helpers.JSONSuccessResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "", category)
}
