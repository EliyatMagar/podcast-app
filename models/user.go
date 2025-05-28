package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"` // e.g "user", "artist", "admin"
	Playlists []Playlist
	Likes     []Like
	Follows   []Follow
	CreatedAt time.Time
	UpdatedAt time.Time
}
