package repositories

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/entities"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	user := entities.User{
		UserID:    "001",
		FirstName: "Taro",
		LastName:  "Yamada",
	}
	storage := map[string]entities.User{user.UserID: user}
	inst := UserRepositoryMemory{storage}
	ret, err := inst.Get("001")
	if err != nil {
		t.Errorf("got = %v, want = %v", err, nil)
	}
	if !reflect.DeepEqual(ret, user) {
		t.Errorf("got = %v, want = %v", ret, user)
	}
}

func TestSave(t *testing.T) {
	storage := map[string]entities.User{}
	inst := UserRepositoryMemory{storage}
	user := entities.User{
		UserID:    "001",
		FirstName: "Taro",
		LastName:  "Yamada",
	}
	err := inst.Save(user, false)
	if err != nil {
		t.Errorf("got = %v, want = %v", err, nil)
	}
	err = inst.Save(user, false)
	if err == nil {
		t.Errorf("got = %v, not want = %v", err, nil)
	}
}
