package routes

import (
	"go-podcast/controllers"
	"go-podcast/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPodcastRoutes(router *gin.Engine) {
	r := router.Group("/api/podcasts")
	r.Use(middleware.AuthMiddleware())
	{
		r.POST("/", controllers.CreatePodcast)
		r.GET("/", controllers.GetAllPodcasts)
		r.GET("/:id", controllers.GetPodcastByID)
		r.PUT("/:id", controllers.UpdatePodcast)
		r.DELETE("/:id", controllers.DeletePodcast)
	}
}
