package domain

import "github.com/jinzhu/gorm"

// Game is infomation of novel games
type Game struct {
	gorm.Model
	GameID    string
	Title     string
	Thumbnail string
	IsR18     bool
	URL       string
}
