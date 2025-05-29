package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateArtist - POST /api/artists
func CreateArtist(c *gin.Context) {
	// Get logged-in user ID from context (assumed set by middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if artist already exists for this user
	var existingArtist models.Artist
	if err := database.DB.Where("user_id = ?", userID).First(&existingArtist).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Artist profile already exists for this user"})
		return
	}

	// Bind input
	var artistInput models.Artist
	if err := c.ShouldBindJSON(&artistInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign userID
	artistInput.UserID = userID.(uint)

	if err := database.DB.Create(&artistInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create artist"})
		return
	}

	c.JSON(http.StatusCreated, artistInput)
}

// GetArtistByUserID - GET /api/artists/me
func GetArtistByUserID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var artist models.Artist
	if err := database.DB.Preload("Podcasts").Preload("Albums").Where("user_id = ?", userID).First(&artist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artist profile not found"})
		return
	}

	c.JSON(http.StatusOK, artist)
}

// GetAllArtists - GET /api/artists
func GetAllArtists(c *gin.Context) {
	var artists []models.Artist
	if err := database.DB.Preload("User").Find(&artists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch artists"})
		return
	}
	c.JSON(http.StatusOK, artists)
}

// GetArtistByID - GET /api/artists/:id
func GetArtistByID(c *gin.Context) {
	id := c.Param("id")
	var artist models.Artist

	if err := database.DB.Preload("User").First(&artist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artist not found"})
		return
	}
	c.JSON(http.StatusOK, artist)
}

// UpdateArtistByID - PUT /api/artists/:id
func UpdateArtistByID(c *gin.Context) {
	id := c.Param("id")
	var artist models.Artist

	if err := database.DB.First(&artist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artist not found"})
		return
	}

	var updatedData models.Artist
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	artist.Name = updatedData.Name
	artist.Bio = updatedData.Bio
	artist.ImageURL = updatedData.ImageURL

	if err := database.DB.Save(&artist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update artist"})
		return
	}
	c.JSON(http.StatusOK, artist)
}

// DeleteArtistByID - DELETE /api/artists/:id
func DeleteArtistByID(c *gin.Context) {
	id := c.Param("id")
	var artist models.Artist

	if err := database.DB.First(&artist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artist not found"})
		return
	}

	if err := database.DB.Delete(&artist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete artist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Artist deleted successfully"})
}
