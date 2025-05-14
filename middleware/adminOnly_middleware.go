package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"net/http"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")

		//fmt.Println("Extracted role from JWT:", role) // Log

		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, web.ApiResponse{
				Code:    "FORBIDDEN",
				Message: "Access Denied: Admins Only",
				Data:    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
