package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
)

type MainController struct {
}

func NewMainController() *MainController {
	return &MainController{}
}

func (h *MainController) MainController(c *gin.Context) {
	helpers.JSONSuccessResponse(c, "Simple Online Shop API is Running", nil)
}
