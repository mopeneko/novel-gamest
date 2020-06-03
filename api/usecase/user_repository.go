package usecase

import (
	"github.com/mopeneko/novel-gamest/api/domain"
)

// UserRepository manages User entities
type UserRepository interface {
	Save(domain.User) error
	FindByID(string) (domain.User, error)
	FindByName(string) ([]domain.User, error)
}
