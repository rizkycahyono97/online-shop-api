package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"net/http"
)

type MainController struct {
}

func NewMainController() *MainController {
	return &MainController{}
}

func (h *MainController) MainController(c *gin.Context) {
	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Simple Online Shop API is Running",
	})
}
