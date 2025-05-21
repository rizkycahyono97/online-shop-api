package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rizkycahyono97/online-shop-api/config"
	"github.com/rizkycahyono97/online-shop-api/controller"
	"github.com/rizkycahyono97/online-shop-api/middleware"
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

	// initialize database
	config.InitDB()

	// main inject
	mainController := controller.NewMainController()

	// auth inject
	authRepo := repositories.NewAuthRepositoryImpl(config.DB)
	authService := services.NewAuthServiceImpl(authRepo)
	authController := controller.NewAuthController(authService)

	// user inject
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// product inject
	productRepo := repositories.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// cart inject
	cartRepo := repositories.NewCartRepository(config.DB)
	cartService := services.NewCartService(cartRepo, productRepo)
	cartController := controller.NewCartController(cartService)

	//  order inject
	orderRepo := repositories.NewOrderRepository(config.DB)
	orderService := services.NewOrderService(orderRepo, cartRepo, productRepo)
	orderController := controller.NewOrderController(orderService)

	// payment inject
	paymentRepo := repositories.NewPaymentRepository(config.DB)
	paymentService := services.NewPaymentService(paymentRepo, orderRepo)
	paymentController := controller.NewPaymentController(paymentService)

	// category inject
	categoryRepo := repositories.NewCategoryRepository(config.DB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	// initialize route
	r := gin.Default()

	// Apply the CORS middleware to the router.
	r.Use(cors.New(middleware.SetupCors()))

	// router
	routes.SetupRoutes(
		r,
		mainController,
		authController,
		userController,
		productController,
		cartController,
		orderController,
		paymentController,
		categoryController)

	// Retrieve the application port from environment variables with a default value of "8080".
	appPort := config.GetEnv("APP_PORT", "8080")

	// run application
	if err := r.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
