package routes

import (
	"go-podcast/controllers"

	"github.com/gin-gonic/gin"
)

func FollowRoutes(router *gin.Engine) {
	f := router.Group("/api/follows")
	{
		f.POST("/artist/:artistID", controllers.FollowArtist)
		f.POST("/podcast/:podcastID", controllers.FollowPodcast)
		f.DELETE("/:id", controllers.Unfollow)
		f.GET("", controllers.GetUserFollows)
	}
}
