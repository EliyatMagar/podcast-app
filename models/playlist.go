package models

import "time"

type Playlist struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	UserID    uint
	User      User
	Tracks    []Track `gorm:"many2many:playlist_tracks"`
	CreatedAt time.Time
}
