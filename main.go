package main

import (
	"log"
	"os"

	"go-podcast/config"
	"go-podcast/database"
	"go-podcast/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	router := gin.Default()

	routes.UserRoutes(router)
	routes.RegisterEpisodeRoutes(router)
	routes.PodcastRoutes(router)
	routes.ArtistRoutes(router)
	routes.AlbumRoutes(router)
	routes.TrackRoutes(router)
	routes.PlaylistRoutes(router)
	routes.LikeRoutes(router)
	routes.FollowRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}
