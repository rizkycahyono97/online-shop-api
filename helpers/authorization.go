package helpers

import "github.com/gin-gonic/gin"

func IsOwnerOrAdmin(c *gin.Context, targetID uint) bool {
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		return false
	}

	role, _ := c.Get("user_role")

	// JWT Claims
	userID := uint(userIDFromToken.(float64))

	// Return true jika user sedang mengakses dirinya sendiri atau admin
	return userID == targetID || role == "admin"
}
