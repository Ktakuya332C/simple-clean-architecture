package usecases

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/entities"
	"github.com/Ktakuya332C/simple-clean-architecture/mock/repositories"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := entities.User{
		UserID:    "001",
		FirstName: "Taro",
		LastName:  "Yamada",
	}

	mockUserRepository := repositories.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Save(user, false).Return(nil)
	createUser := CreateUser{UserRepository: mockUserRepository}
	err := createUser.Create(user)
	if err != nil {
		t.Errorf("%v", err)
	}
}
