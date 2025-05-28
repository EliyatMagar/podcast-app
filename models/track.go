package models

import "time"

type Track struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	AlbumID   uint
	Album     Album
	AudioURL  string
	Duration  int //In Seconds
	Likes     []Like
	CreatedAt time.Time
	UpdatedAt time.Time
}
