package routes

import (
	"go-podcast/controllers"
	"go-podcast/middleware"

	"github.com/gin-gonic/gin"
)

func PodcastRoutes(router *gin.Engine) {
	podcastGroup := router.Group("/api/podcasts")
	{
		podcastGroup.GET("", controllers.GetAllPodcasts)
		podcastGroup.GET("/:id", controllers.GetPodcastByID)

		// Protected routes
		podcastGroup.Use(middleware.AuthMiddleware())
		{
			podcastGroup.POST("", controllers.CreatePodcast)
			podcastGroup.PUT("/:id", controllers.UpdatePodcast)
			podcastGroup.DELETE("/:id", controllers.DeletePodcast)
		}
	}
}
