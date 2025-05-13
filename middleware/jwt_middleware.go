package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rizkycahyono97/online-shop-api/config"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"net/http"
	"strings"
)

var jwtSecret = []byte(config.GetEnv("JWT_SECRET", "super_secret_key_123!@#"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		//fmt.Println("Authorization Header:", authHeader) // Log

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "UNAUTHORIZED",
				Message: "JWT Is Null",
				Data:    nil,
			})
			c.Abort()
			return
		}

		// The token should be prefixed with "Bearer "
		parts := strings.Split(authHeader, " ")

		//fmt.Println("Parts:", parts) // Log

		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println("Invalid Authorization Header Format") // Log
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		//fmt.Println("Token String:", tokenString) // Log

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {

			//fmt.Println("JWT Parse Error:", err) // Log

			c.JSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "UNAUTHORIZED",
				Message: "JWT Parse Error",
				Data:    nil,
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {

			//fmt.Println("Failed to parse claims") // Log

			c.JSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "UNAUTHORIZED",
				Message: "Failed to parse claims",
				Data:    nil,
			})
			c.Abort()
			return
		}

		//fmt.Println("Claims:", claims) // Log

		// Inject user info to context
		c.Set("user_id", claims["user_id"])
		c.Set("user_email", claims["email"])
		c.Set("user_role", claims["role"])

		c.Next()
	}
}
