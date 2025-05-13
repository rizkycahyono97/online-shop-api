package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/controller"
	"github.com/rizkycahyono97/online-shop-api/middleware"
)

func SetupRoutes(
	r *gin.Engine,
	authController *controller.AuthController,
	userController *controller.UserController) {
	// public routes
	public := r.Group("/api/v1")
	{
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
		protected.GET("/users", userController.GetAllProfiles)
		protected.PUT("/users/:id", userController.UpdateProfile)
		protected.DELETE("/users/:id", userController.DeleteProfile)
	}
}
