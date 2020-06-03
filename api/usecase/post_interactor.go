package usecase

import (
	"github.com/mopeneko/novel-gamest/api/domain"
)

// PostInteractor is a bridge to interact Post entities
type PostInteractor struct {
	PostRepository PostRepository
}

// Add post
func (interactor PostInteractor) Add(post domain.Post) error {
	return interactor.PostRepository.Save(post)
}

// GetByID returns post which has same id
func (interactor PostInteractor) GetByID(id string) (domain.Post, error) {
	return interactor.PostRepository.FindByID(id)
}

// GetByGameID return posts which is written for the game
func (interactor PostInteractor) GetByGameID(id string) ([]domain.Post, error) {
	return interactor.PostRepository.FindByGameID(id)
}

// GetByUserID return posts which is written by the user
func (interactor PostInteractor) GetByUserID(id string) ([]domain.Post, error) {
	return interactor.PostRepository.FindByUserID(id)
}
