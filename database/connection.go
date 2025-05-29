package database

import (
	"fmt"
	"go-podcast/config"
	"go-podcast/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Asia/Kathmandu",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.DBName,
		config.DB.Port,
	)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect PostgreSQL:", err)
	}
	log.Println("Connected to PostgreSQL database")

	//Auto migrate models

	if err := DB.AutoMigrate(&models.User{}, &models.Episode{}, &models.Podcast{}, &models.Artist{}, &models.Album{}, &models.Like{}, &models.Playlist{}, &models.Track{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
}
