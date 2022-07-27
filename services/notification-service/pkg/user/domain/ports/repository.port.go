package ports

import (
	"notifications/pkg/user/domain/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	UpdateUser(query interface{}, userId uuid.UUID) error
	CreateUser(user models.User) error
}
