package ports

import (
	"service-users/pkg/user/domain/events"
	"service-users/pkg/user/domain/models"

	"github.com/google/uuid"
)

type UserRepo interface {
	Create(new models.User) (events.Event, bool, error)
	GetPublicUser(uuid uuid.UUID) (models.UserPublic, bool, error)
	GetPrivateUser(document string, value interface{}) (models.User, bool, error)
}
