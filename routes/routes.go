package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/controller"
	"github.com/rizkycahyono97/online-shop-api/middleware"
)

func SetupRoutes(
	r *gin.Engine,
	mainController *controller.MainController,
	authController *controller.AuthController,
	userController *controller.UserController,
	productController *controller.ProductController,
	cartController *controller.CartController) {

	// public routes
	public := r.Group("/api/v1")
	{
		// Main
		public.GET("/", mainController.MainController)

		// auth endpoint
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
	}

	// protected routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		// user endpoint
		protected.GET("/users/:id", userController.GetProfile)
		protected.PUT("/users/:id", userController.UpdateProfile)

		// product endpoint
		protected.GET("/products", productController.GetAllProducts)
		protected.GET("/products/:id", productController.GetProductById)

		// cart endpoint
		protected.GET("/cart/items", cartController.GetCartItems)
		protected.POST("/cart/items", cartController.AddItem)
		protected.DELETE("/cart/items/:product_id", cartController.RemoveItemFromCart)
	}

	// admins routes
	adminRoutes := r.Group("/api/v1")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		//user endpoint
		adminRoutes.GET("/users", userController.GetAllProfiles)
		adminRoutes.DELETE("/users/:id", userController.DeleteProfile)

		// product endpoint
		adminRoutes.POST("/products", productController.CreateProduct)
		adminRoutes.PUT("/products/:id", productController.UpdateProduct)
		adminRoutes.DELETE("/products/:id", productController.DeleteProduct)

		// cart endpoint
		adminRoutes.GET("/cart", cartController.GetAllCarts)
	}
}
