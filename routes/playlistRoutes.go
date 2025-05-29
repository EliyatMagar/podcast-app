package routes

import (
	"go-podcast/controllers"

	"github.com/gin-gonic/gin"
)

func PlaylistRoutes(router *gin.Engine) {
	playlists := router.Group("/api/playlists")
	{
		playlists.POST("", controllers.CreatePlaylist)
		playlists.GET("", controllers.GetAllPlaylists)
		playlists.GET("/:id", controllers.GetPlaylistByID)
		playlists.PUT("/:id/tracks", controllers.UpdatePlaylistTracks)
		playlists.DELETE("/:id", controllers.DeletePlaylist)
	}
}
