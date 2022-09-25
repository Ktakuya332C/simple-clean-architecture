package repositories

import (
	"fmt"
	"github.com/Ktakuya332C/simple-clean-architecture/internal/entities"
)

type UserRepositoryMemory struct {
	storage map[string]entities.User
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	storage := map[string]entities.User{}
	return &UserRepositoryMemory{storage}
}

func (inst UserRepositoryMemory) Get(userID string) (entities.User, error) {
	if user, ok := inst.storage[userID]; ok {
		return user, nil
	}
	return entities.User{}, fmt.Errorf("Not found %v", userID)
}

func (inst *UserRepositoryMemory) Save(user entities.User, overwrite bool) error {
	if _, found := inst.storage[user.UserID]; found {
		return fmt.Errorf("Found %v", user.UserID)
	}
	inst.storage[user.UserID] = user
	return nil
}
