package domain

import (
	"github.com/jinzhu/gorm"
)

// Post of the service
type Post struct {
	gorm.Model
	PostID string
	Text   string
	Author User
	Game   Game
}
