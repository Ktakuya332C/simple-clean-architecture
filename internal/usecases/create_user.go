package usecases

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/entities"
	"github.com/Ktakuya332C/simple-clean-architecture/internal/repositories"
)

type CreateUser struct {
	UserRepository repositories.UserRepository
}

func (inst *CreateUser) Create(user entities.User) error {
	return inst.UserRepository.Save(user, false)
}
