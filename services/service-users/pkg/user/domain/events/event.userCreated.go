package events

import (
	"service-users/pkg/user/domain/models"

	"github.com/google/uuid"
)

type UserCreated struct {
	UserId     uuid.UUID   `json:"id"`
	UserName   models.Name `json:"name"`
	UserAvatar string      `json:"avatar"`
}

func (u UserCreated) Name() string {
	return "event.user.created"
}
func (u UserCreated) Exchange() string {
	return "User"
}
func (u UserCreated) Routing() string {
	return "created"
}
