package database

import (
	"github.com/jinzhu/gorm"
	"github.com/mopeneko/novel-gamest/api/domain"
)

// UserRepository implements usercase.UserRepository
type UserRepository struct {
	DB *gorm.DB
}

// Save an user information
func (repo *UserRepository) Save(user domain.User) error {
	return repo.DB.Create(&user).Error
}

// FindByID returns an user which has same ID
func (repo *UserRepository) FindByID(id string) (domain.User, error) {
	user := domain.User{
		UserID: id,
	}
	err := repo.DB.First(&user).Error

	return user, err
}
