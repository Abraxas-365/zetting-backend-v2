package events

import "github.com/google/uuid"

type UserCreated struct {
	UserId   uuid.UUID `json:"user_id"`
	UserName string    `json:"user_name"`
}

func (u UserCreated) Name() string {
	return "event.user.created"
}
func (u UserCreated) Exchange() string {
	return "User"
}
func (u UserCreated) Routing() string {
	return "creted"
}
