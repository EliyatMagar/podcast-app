package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// POST /api/likes/track/:trackID
func LikeTrack(c *gin.Context) {
	userID, _ := c.Get("userID")
	trackIDStr := c.Param("trackID")
	trackID, err := strconv.ParseUint(trackIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track ID"})
		return
	}

	like := models.Like{
		UserID:  userID.(uint),
		TrackID: uintPtr(uint(trackID)),
	}

	if err := database.DB.Create(&like).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like track"})
		return
	}

	c.JSON(http.StatusOK, like)
}

// POST /api/likes/episode/:episodeID
func LikeEpisode(c *gin.Context) {
	userID, _ := c.Get("userID")
	episodeIDStr := c.Param("episodeID")
	episodeID, err := strconv.ParseUint(episodeIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid episode ID"})
		return
	}

	like := models.Like{
		UserID:    userID.(uint),
		EpisodeID: uintPtr(uint(episodeID)),
	}

	if err := database.DB.Create(&like).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like episode"})
		return
	}

	c.JSON(http.StatusOK, like)
}

// DELETE /api/likes/:id
func Unlike(c *gin.Context) {
	id := c.Param("id")
	var like models.Like

	if err := database.DB.First(&like, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Like not found"})
		return
	}

	if err := database.DB.Delete(&like).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like removed"})
}

// GET /api/likes?user=1
func GetUserLikes(c *gin.Context) {
	userIDStr := c.Query("user")
	var likes []models.Like

	query := database.DB
	if userIDStr != "" {
		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&likes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch likes"})
		return
	}

	c.JSON(http.StatusOK, likes)
}
