package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"net/http"
	"strings"
)

var jwtSecret = []byte("secret")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.GetHeader("Authorization")
		if authHandler == "" {
			c.JSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "UNAUTHORIZED",
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    nil,
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHandler, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "UNAUTHORIZED",
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    nil,
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "UNAUTHORIZED",
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    nil,
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "UNAUTHORIZED",
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    nil,
			})
		}

		// Inject user info to context
		c.Set("user_id", claims["user_id"])
		c.Set("user_email", claims["email"])
		c.Set("user_role", claims["role"])

		c.Next()
	}
}
