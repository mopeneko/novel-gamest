package domain

import (
	"github.com/jinzhu/gorm"
)

// User of the service
type User struct {
	gorm.Model
	UserID   string
	Password string
	Name     string
}
