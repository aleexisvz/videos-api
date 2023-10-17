package models

import (
	"time"

	"gorm.io/gorm"
)

// Video struct
type Video struct {
	gorm.Model
	ID        uint
	Title     string
	URL       string
	Author    string
	Views     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName returns the name of the table
func (Video) TableName() string {
	return "videos"
}
