package models

import "time"

type Album struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"not null"`
	ArtistID   uint
	Artist     Artist
	CoverURL   string
	Tracks     []Track
	ReleasedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
