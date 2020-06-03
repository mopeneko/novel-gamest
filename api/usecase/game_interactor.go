package usecase

import (
	"github.com/mopeneko/novel-gamest/api/domain"
)

// GameInteractor is a bridge to interact Game entities
type GameInteractor struct {
	GameRepository GameRepository
}

// Add game
func (interactor *GameInteractor) Add(game domain.Game) error {
	return interactor.GameRepository.Save(game)
}

// GetByID returns game which has same id
func (interactor *GameInteractor) GetByID(id string) (domain.Game, error) {
	return interactor.GameRepository.FindByID(id)
}

// GetByTitle returns games which has same title
func (interactor *GameInteractor) GetByTitle(title string) ([]domain.Game, error) {
	return interactor.GameRepository.FindByTitle(title)
}
