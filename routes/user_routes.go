package routes

import (
	"go-podcast/controllers"
	"go-podcast/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	public := r.Group("/api/user")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("api/user")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.Profile)
		protected.PUT("/profileUpdate", controllers.UpdateProfile)
		protected.POST("/logout", controllers.Logout)
	}
}
