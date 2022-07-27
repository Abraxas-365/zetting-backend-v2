package models

import (
	"notifications/pkg/user/domain/models"
	"time"

	"github.com/google/uuid"
)

type NotificationInput struct {
	ID           uuid.UUID `bson:"_id,omitempty" json:"id"`
	Reciver      uuid.UUID `bson:"reciver" json:"reciver"` //user who resive the data
	Emitter      uuid.UUID `bson:"emitter" json:"emitter"` //user who send the data
	Kind         string    `bson:"kind" json:"kind"`
	CreationDate time.Time `bson:"creation_date" json:"creation_date"`
	Status       bool      `bson:"status" json:"status"` //true if the notification have been seen
}

type NotificationQuery struct {
	ID           uuid.UUID   `bson:"_id,omitempty" json:"id"`
	Reciver      models.User `bson:"reciver" json:"reciver"` //user who resive the data
	Emitter      models.User `bson:"emitter" json:"emitter"` //user who send the data
	Kind         string      `bson:"kind" json:"kind"`
	CreationDate time.Time   `bson:"creation_date" json:"creation_date"`
	Status       bool        `bson:"status" json:"status"` //true if the notification have been seen
}

type NotificationsQuery []*NotificationQuery

func (n *NotificationInput) New() *NotificationInput {
	n.ID = uuid.New()
	n.CreationDate = time.Now()

	return n
}
