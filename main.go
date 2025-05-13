package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rizkycahyono97/online-shop-api/config"
	"github.com/rizkycahyono97/online-shop-api/controller"
	"github.com/rizkycahyono97/online-shop-api/repositories"
	"github.com/rizkycahyono97/online-shop-api/routes"
	"github.com/rizkycahyono97/online-shop-api/services"
	"log"
)

func main() {
	//Load .env File
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.InitDB()

	authRepo := repositories.NewAuthRepositoryImpl(config.DB)
	authService := services.NewAuthServiceImpl(authRepo)
	authController := controller.NewAuthController(authService)

	r := gin.Default()

	routes.SetupRoutes(r, authController)

	r.Run("localhost:8080")
}
