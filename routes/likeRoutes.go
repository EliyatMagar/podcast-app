package routes

import (
	"go-podcast/controllers"

	"github.com/gin-gonic/gin"
)

func LikeRoutes(router *gin.Engine) {
	likes := router.Group("/api/likes")
	{
		likes.POST("/track/:trackID", controllers.LikeTrack)
		likes.POST("/episode/:episodeID", controllers.LikeEpisode)
		likes.DELETE("/:id", controllers.Unlike)
		likes.GET("", controllers.GetUserLikes)
	}
}
