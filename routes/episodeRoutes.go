package routes

import (
	"go-podcast/controllers"
	"go-podcast/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEpisodeRoutes(router *gin.Engine) {
	episode := router.Group("/api/episodes")
	episode.Use(middleware.AuthMiddleware())

	episode.POST("/", controllers.CreateEpisode)
	episode.GET("/", controllers.GetEpisodes)
	episode.GET("/:id", controllers.GetEpisodeByID)
	episode.PUT("/:id", controllers.UpdateEpisodeByID)
	episode.DELETE("/:id", controllers.DeleteEpisodeByID)
}
