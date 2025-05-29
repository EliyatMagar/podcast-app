package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePlaylist - POST /api/playlists
func CreatePlaylist(c *gin.Context) {
	var playlist models.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&playlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create playlist"})
		return
	}

	c.JSON(http.StatusCreated, playlist)
}

// GetAllPlaylists - GET /api/playlists
func GetAllPlaylists(c *gin.Context) {
	var playlists []models.Playlist
	if err := database.DB.Preload("Tracks").Find(&playlists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists"})
		return
	}
	c.JSON(http.StatusOK, playlists)
}

// GetPlaylistByID - GET /api/playlists/:id
func GetPlaylistByID(c *gin.Context) {
	id := c.Param("id")
	var playlist models.Playlist

	if err := database.DB.Preload("Tracks").First(&playlist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}
	c.JSON(http.StatusOK, playlist)
}

// UpdatePlaylistTracks - PUT /api/playlists/:id/tracks
func UpdatePlaylistTracks(c *gin.Context) {
	id := c.Param("id")
	var playlist models.Playlist

	if err := database.DB.First(&playlist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	var trackIDs []uint
	if err := c.ShouldBindJSON(&trackIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track IDs"})
		return
	}

	var tracks []models.Track
	if err := database.DB.Where("id IN ?", trackIDs).Find(&tracks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tracks"})
		return
	}

	if err := database.DB.Model(&playlist).Association("Tracks").Replace(tracks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update playlist tracks"})
		return
	}

	c.JSON(http.StatusOK, playlist)
}

// DeletePlaylist - DELETE /api/playlists/:id
func DeletePlaylist(c *gin.Context) {
	id := c.Param("id")
	var playlist models.Playlist

	if err := database.DB.First(&playlist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	if err := database.DB.Delete(&playlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete playlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist deleted successfully"})
}
