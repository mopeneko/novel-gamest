package usecase

import (
	"github.com/mopeneko/novel-gamest/api/domain"
)

// UserInteractor is a bridge to interact User entities
type UserInteractor struct {
	UserRepository UserRepository
}

// Add user
func (interactor UserInteractor) Add(user domain.User) error {
	return interactor.UserRepository.Save(user)
}

// GetByID returns user which has same id
func (interactor UserInteractor) GetByID(id string) (domain.User, error) {
	return interactor.UserRepository.FindByID(id)
}

// GetByName returns users which has same name
func (interactor *UserInteractor) GetByName(name string) ([]domain.User, error) {
	return interactor.UserRepository.FindByName(name)
}
