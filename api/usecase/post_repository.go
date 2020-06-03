package usecase

import (
	"github.com/mopeneko/novel-gamest/api/domain"
)

// PostRepository manages Post entities
type PostRepository interface {
	Save(domain.Post) error
	FindByID(string) (domain.Post, error)
	FindByGameID(string) ([]domain.Post, error)
	FindByUserID(string) ([]domain.Post, error)
}
