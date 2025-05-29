package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePodcast - POST /api/podcasts
func CreatePodcast(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if user is an artist
	var artist models.Artist
	if err := database.DB.Where("user_id = ?", userID).First(&artist).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Artist profile not found"})
		return
	}

	// Bind request body
	var input models.Podcast
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign artist ID
	input.ArtistID = artist.ID

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create podcast"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GetAllPodcasts - GET /api/podcasts
func GetAllPodcasts(c *gin.Context) {
	var podcasts []models.Podcast
	if err := database.DB.Preload("Artist").Find(&podcasts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch podcasts"})
		return
	}
	c.JSON(http.StatusOK, podcasts)
}

// GetPodcastByID - GET /api/podcasts/:id
func GetPodcastByID(c *gin.Context) {
	id := c.Param("id")
	var podcast models.Podcast

	if err := database.DB.Preload("Artist").Preload("Episodes").First(&podcast, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Podcast not found"})
		return
	}

	c.JSON(http.StatusOK, podcast)
}

// UpdatePodcast - PUT /api/podcasts/:id
func UpdatePodcast(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var podcast models.Podcast
	if err := database.DB.First(&podcast, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Podcast not found"})
		return
	}

	// Check ownership
	var artist models.Artist
	if err := database.DB.First(&artist, "user_id = ?", userID).Error; err != nil || artist.ID != podcast.ArtistID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this podcast"})
		return
	}

	var updatedData models.Podcast
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	podcast.Title = updatedData.Title
	podcast.Description = updatedData.Description
	podcast.CoverURL = updatedData.CoverURL
	podcast.Category = updatedData.Category

	if err := database.DB.Save(&podcast).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update podcast"})
		return
	}

	c.JSON(http.StatusOK, podcast)
}

// DeletePodcast - DELETE /api/podcasts/:id
func DeletePodcast(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var podcast models.Podcast
	if err := database.DB.First(&podcast, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Podcast not found"})
		return
	}

	// Check ownership
	var artist models.Artist
	if err := database.DB.First(&artist, "user_id = ?", userID).Error; err != nil || artist.ID != podcast.ArtistID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this podcast"})
		return
	}

	if err := database.DB.Delete(&podcast).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete podcast"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Podcast deleted successfully"})
}
