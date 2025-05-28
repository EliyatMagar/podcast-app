package models

import "time"

type Podcast struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	CoverURL    string
	ArtistID    uint
	Artist      Artist
	Episodes    []Episode
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
