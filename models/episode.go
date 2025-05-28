package models

import "time"

type Episode struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	PodcastID   uint
	Podcast     Podcast
	AudioURL    string
	Description string
	Duration    int //in seconds
	ReleasedAt  time.Time
}
