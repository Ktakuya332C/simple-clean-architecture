package usecases

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/entities"
	"github.com/Ktakuya332C/simple-clean-architecture/internal/repositories"
)

type GetUser struct {
	UserRepository repositories.UserRepository
}

func (inst GetUser) Get(userID string) (entities.User, error) {
	return inst.UserRepository.Get(userID)
}
