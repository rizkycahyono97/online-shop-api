package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// register handler
func (ac *AuthController) Register(c *gin.Context) {
	var input web.UserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	user, err := ac.authService.Register(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, web.ApiResponse{
		Code:    "created",
		Message: "User Registered Successfully",
		Data:    web.UserResponseFromModel(user),
	})
}

// Login Handler
func (ac *AuthController) Login(c *gin.Context) {
	var input web.UserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	token, err := ac.authService.Login(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Login Successfully",
		Token:   token,
	})
}
