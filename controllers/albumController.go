package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAlbum - POST /api/albums
func CreateAlbum(c *gin.Context) {
	userID, _ := c.Get("userID")

	// Find the artist based on userID
	var artist models.Artist
	if err := database.DB.Where("user_id = ?", userID).First(&artist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Artist profile not found for this user"})
		return
	}

	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album.ArtistID = artist.ID

	if err := database.DB.Create(&album).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(http.StatusCreated, album)
}

// GetAllAlbums - GET /api/albums
func GetAllAlbums(c *gin.Context) {
	var albums []models.Album
	if err := database.DB.Preload("Artist").Find(&albums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch albums"})
		return
	}
	c.JSON(http.StatusOK, albums)
}

// GetAlbumByID - GET /api/albums/:id
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := database.DB.Preload("Tracks").Preload("Artist").First(&album, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}
	c.JSON(http.StatusOK, album)
}

// UpdateAlbumByID - PUT /api/albums/:id
func UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := database.DB.First(&album, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	var updatedData models.Album
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album.Title = updatedData.Title
	album.CoverURL = updatedData.CoverURL
	album.ReleasedAt = updatedData.ReleasedAt

	if err := database.DB.Save(&album).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album"})
		return
	}

	c.JSON(http.StatusOK, album)
}

// DeleteAlbumByID - DELETE /api/albums/:id
func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := database.DB.First(&album, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	if err := database.DB.Delete(&album).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
