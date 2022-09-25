//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repositories

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/entities"
)

type UserRepository interface {
	Save(user entities.User, overwrite bool) error
	Get(userID string) (entities.User, error)
}
