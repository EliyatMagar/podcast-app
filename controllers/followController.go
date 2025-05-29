package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Helper: convert uint to pointer
func uintPtr(i uint) *uint {
	return &i
}

// POST /api/follows/artist/:artistID
func FollowArtist(c *gin.Context) {
	userID, _ := c.Get("userID")
	artistID, err := strconv.ParseUint(c.Param("artistID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid artist ID"})
		return
	}

	follow := models.Follow{
		UserID:   userID.(uint),
		ArtistID: uintPtr(uint(artistID)),
	}

	if err := database.DB.Create(&follow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow artist"})
		return
	}

	c.JSON(http.StatusOK, follow)
}

// POST /api/follows/podcast/:podcastID
func FollowPodcast(c *gin.Context) {
	userID, _ := c.Get("userID")
	podcastID, err := strconv.ParseUint(c.Param("podcastID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid podcast ID"})
		return
	}

	follow := models.Follow{
		UserID:    userID.(uint),
		PodcastID: uintPtr(uint(podcastID)),
	}

	if err := database.DB.Create(&follow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow podcast"})
		return
	}

	c.JSON(http.StatusOK, follow)
}

// DELETE /api/follows/:id
func Unfollow(c *gin.Context) {
	followID := c.Param("id")

	var follow models.Follow
	if err := database.DB.First(&follow, followID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Follow not found"})
		return
	}

	if err := database.DB.Delete(&follow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unfollowed successfully"})
}

// GET /api/follows?user=1
func GetUserFollows(c *gin.Context) {
	userIDStr := c.Query("user")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var follows []models.Follow
	if err := database.DB.Where("user_id = ?", userID).Find(&follows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get follows"})
		return
	}

	c.JSON(http.StatusOK, follows)
}
