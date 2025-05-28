package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePodcast handles POST /api/podcasts
func CreatePodcast(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "artist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only artists can create podcasts"})
		return
	}
	userID, _ := c.Get("userID")

	var input models.Podcast
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ArtistID = userID.(uint)
	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create podcast"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func GetAllPodcasts(c *gin.Context) {
	var podcasts []models.Podcast
	database.DB.Preload("Episodes").Find(&podcasts)
	c.JSON(http.StatusOK, podcasts)
}

func GetPodcastByID(c *gin.Context) {
	var podcast models.Podcast
	id := c.Param("id")
	if err := database.DB.Preload("Episodes").First(&podcast, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Podcast not found"})
		return
	}
	c.JSON(http.StatusOK, podcast)
}

func UpdatePodcast(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var podcast models.Podcast
	if err := database.DB.First(&podcast, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Podcast not found"})
		return
	}

	if podcast.ArtistID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this podcast"})
		return
	}

	var input models.Podcast
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	podcast.Title = input.Title
	podcast.Description = input.Description
	podcast.CoverURL = input.CoverURL
	podcast.Category = input.Category

	database.DB.Save(&podcast)
	c.JSON(http.StatusOK, podcast)
}

func DeletePodcast(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var podcast models.Podcast
	if err := database.DB.First(&podcast, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Podcast not found"})
		return
	}

	if podcast.ArtistID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this podcast"})
		return
	}

	database.DB.Delete(&podcast)
	c.JSON(http.StatusOK, gin.H{"message": "Podcast deleted successfully"})
}
