package controllers

import (
	"go-podcast/database"
	"go-podcast/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create Episode
func CreateEpisode(c *gin.Context) {
	var input struct {
		Title       string `json:"title"`
		PodcastID   uint   `json:"podcast_id"`
		AudioURL    string `json:"audio_url"`
		Description string `json:"description"`
		Duration    int    `json:"duration"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	episode := models.Episode{
		Title:       input.Title,
		PodcastID:   input.PodcastID,
		AudioURL:    input.AudioURL,
		Description: input.Description,
		Duration:    input.Duration,
	}

	if err := database.DB.Create(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create episode"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": "Episode created", "episode": episode})
}

//Get All Episodes

func GetEpisodes(c *gin.Context) {
	var episodes []models.Episode
	if err := database.DB.Preload("Podcast").Find(&episodes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch episode"})
		return
	}
	c.JSON(http.StatusOK, episodes)
}

//Get Single Episode by ID

func GetEpisodeByID(c *gin.Context) {
	id := c.Param("id")
	var episode models.Episode

	if err := database.DB.Preload("Podcast").First(&episode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Episode not found"})
		return
	}
	c.JSON(http.StatusOK, episode)
}

//Update Episode

func UpdateEpisodeByID(c *gin.Context) {
	id := c.Param("id")
	var episode models.Episode

	if err := database.DB.First(&episode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Episode not found"})
		return
	}

	var input struct {
		Title       string    `json:"title"`
		AudioURL    string    `json:"audio_url"`
		Description string    `json:"description"`
		Duration    int       `json:"duration"`
		ReleasedAt  time.Time `json:"released_at"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	episode.Title = input.Title
	episode.AudioURL = input.AudioURL
	episode.Description = input.Description
	episode.Duration = input.Duration
	episode.ReleasedAt = input.ReleasedAt

	if err := database.DB.Save(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update episode"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Episode updated", "episode": episode})
}

//Delete Episode

func DeleteEpisodeByID(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Episode{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete episode"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Episode deleted"})
}
