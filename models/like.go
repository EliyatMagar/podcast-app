package models

import "time"

type Like struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	TrackID   *uint // Nullable
	EpisodeID *uint // Nullable
	CreatedAt time.Time
}
