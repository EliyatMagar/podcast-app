package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTrack - POST /api/tracks
func CreateTrack(c *gin.Context) {
	var track models.Track
	if err := c.ShouldBindJSON(&track); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: Validate Album exists
	var album models.Album
	if err := database.DB.First(&album, track.AlbumID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found"})
		return
	}

	if err := database.DB.Create(&track).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create track"})
		return
	}

	c.JSON(http.StatusCreated, track)
}

// GetAllTracks - GET /api/tracks
func GetAllTracks(c *gin.Context) {
	var tracks []models.Track
	if err := database.DB.Preload("Album").Find(&tracks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tracks"})
		return
	}
	c.JSON(http.StatusOK, tracks)
}

// GetTrackByID - GET /api/tracks/:id
func GetTrackByID(c *gin.Context) {
	id := c.Param("id")
	var track models.Track

	if err := database.DB.Preload("Album").First(&track, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}
	c.JSON(http.StatusOK, track)
}

// UpdateTrackByID - PUT /api/tracks/:id
func UpdateTrackByID(c *gin.Context) {
	id := c.Param("id")
	var track models.Track

	if err := database.DB.First(&track, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}

	var updatedData models.Track
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track.Title = updatedData.Title
	track.AudioURL = updatedData.AudioURL
	track.Duration = updatedData.Duration
	track.AlbumID = updatedData.AlbumID

	if err := database.DB.Save(&track).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update track"})
		return
	}

	c.JSON(http.StatusOK, track)
}

// DeleteTrackByID - DELETE /api/tracks/:id
func DeleteTrackByID(c *gin.Context) {
	id := c.Param("id")
	var track models.Track

	if err := database.DB.First(&track, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}

	if err := database.DB.Delete(&track).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete track"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Track deleted successfully"})
}

// GetTracksByAlbumID - GET /api/albums/:albumID/tracks
func GetTracksByAlbumID(c *gin.Context) {
	albumID := c.Param("albumID")
	id, err := strconv.ParseUint(albumID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	var tracks []models.Track
	if err := database.DB.Where("album_id = ?", id).Find(&tracks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tracks for album"})
		return
	}

	c.JSON(http.StatusOK, tracks)
}
