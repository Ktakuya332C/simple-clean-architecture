package usecases

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/entities"
	"github.com/Ktakuya332C/simple-clean-architecture/mock/repositories"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := entities.User{
		UserID:    "001",
		FirstName: "Taro",
		LastName:  "Yamada",
	}

	mockUserRepository := repositories.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Get("001").Return(expected, nil)
	getUser := GetUser{UserRepository: mockUserRepository}
	actual, err := getUser.Get("001")
	if err != nil {
		t.Errorf("got = %v, want = %v", err, nil)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got = %v, want = %v", actual, expected)
	}
}
