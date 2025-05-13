package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
	"net/http"
	"strconv"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// Get Profile by id
func (u *UserController) GetProfile(c *gin.Context) {
	userID := c.GetInt("user_id")

	user, err := u.userService.GetProfile(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, web.ApiResponse{
			Code:    "NOT_FOUND",
			Message: "User not Found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "OK",
		Message: "User Profile Fetched Successfully",
		Data:    user,
	})
}

// Get all profiles
func (u *UserController) GetAllProfiles(c *gin.Context) {
	users, err := u.userService.GetAllProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "SERVER_ERROR",
			Message: "Failed to get Users",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "OK",
		Message: "All Users Fetched Successfully",
		Data:    users,
	})
}

// Update Profile
func (u *UserController) UpdateProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid User ID",
			Data:    nil,
		})
		return
	}

	var req web.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Use the service layer to update the user
	user, err := u.userService.UpdateProfile(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "User Profile Updated Successfully",
		Data:    user,
	})
}

// Delete Profile
func (u *UserController) DeleteProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid User ID",
			Data:    nil,
		})
		return
	}

	err = u.userService.DeleteProfile(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "User Profile Deleted Successfully",
		Data:    nil,
	})
}
