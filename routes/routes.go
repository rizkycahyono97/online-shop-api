package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/controller"
	"github.com/rizkycahyono97/online-shop-api/middleware"
)

func SetupRoutes(
	r *gin.Engine,
	authController *controller.AuthController,
	userController *controller.UserController,
	productController *controller.ProductController) {

	// public routes
	public := r.Group("/api/v1")
	{
		// auth endpoint
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
	}

	// protected routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/")

		// user endpoint
		protected.GET("/users/:id", userController.GetProfile)
		protected.PUT("/users/:id", userController.UpdateProfile)

		//	product endpoint
		protected.GET("/products", productController.GetAllProducts)
		protected.GET("/products/:id", productController.GetProductById)
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
	}
}
