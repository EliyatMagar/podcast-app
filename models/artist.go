package models

import (
	"time"
)

type Artist struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"unique;not null"`
	User      User   `gorm:"foreignKey:UserID"`
	Name      string `gorm:"not null"`
	Bio       string
	ImageURL  string
	Albums    []Album
	Podcasts  []Podcast
	CreatedAt time.Time
	UpdatedAt time.Time
}
