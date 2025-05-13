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

	//AUTHORIZATION
	//ambil user_id dari JWT
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, web.ApiResponse{
			Code:    "UNAUTHORIZED",
			Message: "UNAUTHORIZED",
		})
		return
	}

	// Konversi dari float64 (default MapClaims)
	userID := uint(userIDFromToken.(float64))

	//hanya akses ke dirinya sendiri
	if userID != uint(id) {
		c.JSON(http.StatusForbidden, web.ApiResponse{
			Code:    "FORBIDDEN",
			Message: "You are not allowed to access this user",
		})
		return
	}

	user, err := u.userService.GetProfile(uint(id))
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
		Message: "User Found",
		Data:    web.UserResponseFromModel(user),
	})
}

// Get all profiles
func (u *UserController) GetAllProfiles(c *gin.Context) {
	//ambil role dari JWT
	userRole, exist := c.Get("user_role")
	if !exist {
		c.JSON(http.StatusUnauthorized, web.ApiResponse{
			Code:    "UNAUTHORIZED",
			Message: "Unauthorized",
		})
		return
	}

	role, ok := userRole.(string)
	if !ok || role != "admin" {
		c.JSON(http.StatusUnauthorized, web.ApiResponse{
			Code:    "UNAUTHORIZED",
			Message: "Unauthorized",
			Data:    nil,
		})
		return
	}

	// Jika admin, lanjut ambil data semua user
	users, err := u.userService.GetAllProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "SERVER_ERROR",
			Message: "Failed to get users",
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

	// Ambil user_id dan role dari JWT (di-set oleh AuthMiddleware)
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, web.ApiResponse{
			Code:    "UNAUTHORIZED",
			Message: "Unauthorized",
			Data:    nil,
		})
		return
	}

	role, _ := c.Get("user_role")
	userID := uint(userIDFromToken.(float64))

	// Cek apakah dia user yang bersangkutan atau admin
	if userID != uint(id) && role != "admin" {
		c.JSON(http.StatusForbidden, web.ApiResponse{
			Code:    "FORBIDDEN",
			Message: "You are not allowed to access this user",
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

	// Ambil user_id dan role dari JWT (dari AuthMiddleware)
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, web.ApiResponse{
			Code:    "UNAUTHORIZED",
			Message: "Unauthorized",
			Data:    nil,
		})
		return
	}

	role, _ := c.Get("user_role")
	userID := uint(userIDFromToken.(float64))

	// Hanya user itu sendiri atau admin yang boleh hapus
	if userID != uint(id) && role != "admin" {
		c.JSON(http.StatusForbidden, web.ApiResponse{
			Code:    "FORBIDDEN",
			Message: "You are not allowed to access this user",
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
