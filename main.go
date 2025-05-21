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

	r := gin.Default()

	routes.SetupRoutes(r, mainController, authController, userController, productController, cartController, orderController, paymentController, categoryController)

	r.Run("localhost:8080")
}
