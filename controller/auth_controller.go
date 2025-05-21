package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
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
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	user, err := ac.authService.Register(&input)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "User Registered Successfully", web.UserResponseFromModel(user))
}

// Login Handler
func (ac *AuthController) Login(c *gin.Context) {
	var input web.UserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	token, err := ac.authService.Login(&input)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "Login Successfully", token)
}
