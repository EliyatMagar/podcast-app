package models

import "time"

type Artist struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Bio       string
	ImageURL  string
	Albums    []Album
	Podcasts  []Podcast
	CreatedAt time.Time
	UpdatedAt time.Time
}
