package routes

import (
	"go-podcast/controllers"
	"go-podcast/middleware"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(r *gin.Engine) {
	albumGroup := r.Group("/api/albums")
	{
		albumGroup.GET("", controllers.GetAllAlbums)
		albumGroup.GET("/:id", controllers.GetAlbumByID)

		// Protected routes
		albumGroup.POST("", middleware.AuthMiddleware(), controllers.CreateAlbum)
		albumGroup.PUT("/:id", middleware.AuthMiddleware(), controllers.UpdateAlbumByID)
		albumGroup.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteAlbumByID)
	}
}
