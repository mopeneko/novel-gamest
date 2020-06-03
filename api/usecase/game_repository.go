package usecase

import "github.com/mopeneko/novel-gamest/api/domain"

// GameRepository manages Game entities
type GameRepository interface {
	Save(domain.Game) error
	FindByID(id string) (domain.Game, error)
	FindByTitle(title string) ([]domain.Game, error)
}
