package models

import "time"

type Follow struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ArtistID  *uint // Nullable
	PodcastID *uint // Nullable
	CreatedAt time.Time
}
