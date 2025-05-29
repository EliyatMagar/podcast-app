package routes

import (
	"go-podcast/controllers"
	"go-podcast/middleware"

	"github.com/gin-gonic/gin"
)

func TrackRoutes(r *gin.Engine) {
	trackGroup := r.Group("/api/tracks")
	{
		trackGroup.GET("", controllers.GetAllTracks)
		trackGroup.GET("/:id", controllers.GetTrackByID)

		// Protected routes
		trackGroup.POST("", middleware.AuthMiddleware(), controllers.CreateTrack)
		trackGroup.PUT("/:id", middleware.AuthMiddleware(), controllers.UpdateTrackByID)
		trackGroup.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteTrackByID)
	}

	// Nested route to get tracks by album
	r.GET("/api/albums/:albumID/tracks", controllers.GetTracksByAlbumID)
}
