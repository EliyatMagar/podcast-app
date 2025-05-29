package routes

import (
	"go-podcast/controllers"
	"go-podcast/middleware"

	"github.com/gin-gonic/gin"
)

func ArtistRoutes(router *gin.Engine) {
	artistGroup := router.Group("/api/artists")
	{
		artistGroup.GET("", controllers.GetAllArtists)
		artistGroup.GET("/:id", controllers.GetArtistByID)

		// Auth-protected routes
		artistGroup.Use(middleware.AuthMiddleware()) // Assumes you have middleware to set userID
		{
			artistGroup.POST("", controllers.CreateArtist)
			artistGroup.GET("/me", controllers.GetArtistByUserID)
			artistGroup.PUT("/:id", controllers.UpdateArtistByID)
			artistGroup.DELETE("/:id", controllers.DeleteArtistByID)
		}
	}
}
