package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/helpers"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"github.com/rizkycahyono97/online-shop-api/services"
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
		helpers.JSONBadRequestResponse(c, "Invalid User ID", nil)
		return
	}

	//reusable with helper
	if !helpers.IsOwnerOrAdmin(c, uint(id)) {
		helpers.JSONForbiddenResponse(c, "You are not allowed to update this user", nil)
		return
	}

	user, err := u.userService.GetProfile(uint(id))
	if err != nil {
		helpers.JSONNotFoundResponse(c, "User not found", nil)
		return
	}

	helpers.JSONSuccessResponse(c, "User Found", web.UserResponseFromModel(user))
}

// Get all profiles
func (u *UserController) GetAllProfiles(c *gin.Context) {
	//ambil role dari JWT
	userRole, exist := c.Get("user_role")
	if !exist {
		helpers.JSONForbiddenResponse(c, "Unauthorized", nil)
		return
	}

	role, ok := userRole.(string)
	if !ok || role != "admin" {
		helpers.JSONForbiddenResponse(c, "You are not allowed to access this resource", nil)
		return
	}

	// Jika admin, lanjut ambil data semua user
	users, err := u.userService.GetAllProfiles()
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "Failed to get users", nil)
		return
	}

	helpers.JSONSuccessResponse(c, "All Users Fetched Successfully", users)
}

// Update Profile
func (u *UserController) UpdateProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid User ID", nil)
		return
	}

	//reusable with helper
	if !helpers.IsOwnerOrAdmin(c, uint(id)) {
		helpers.JSONForbiddenResponse(c, "You are not allowed to update this user", nil)
		return
	}

	var req web.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.JSONBadRequestResponse(c, "", err)
	}

	// Use the service layer to update the user
	user, err := u.userService.UpdateProfile(uint(id), &req)
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
	}

	helpers.JSONSuccessResponse(c, "User Profile Updated Successfully", web.UserResponseFromModel(user))
}

// Delete Profile
func (u *UserController) DeleteProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.JSONBadRequestResponse(c, "Invalid User ID", nil)
		return
	}

	// reusable, menggunakan helpers
	if !helpers.IsOwnerOrAdmin(c, uint(id)) {
		helpers.JSONForbiddenResponse(c, "You are not allowed to update this user", nil)
		return
	}

	err = u.userService.DeleteProfile(uint(id))
	if err != nil {
		helpers.JSONInternalErrorResponse(c, "", err)
		return
	}

	helpers.JSONSuccessResponse(c, "User Profile Deleted Successfully", nil)
}
