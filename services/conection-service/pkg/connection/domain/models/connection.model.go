package models

import (
	"connection-service/pkg/user/domain/models"
	"time"

	"github.com/google/uuid"
)

type ConnectionInput struct {
	ID           uuid.UUID `bson:"_id,omitempty" json:"id"`
	UserOne      uuid.UUID `bson:"user_one" json:"user_one"`
	UserTwo      uuid.UUID `bson:"user_two" json:"user_two"`
	CreationDate time.Time `bson:"creation_date" json:"creation_date"`
}

type ConnectionQuery struct {
	ID           uuid.UUID   `bson:"_id,omitempty" json:"id"`
	UserOne      models.User `bson:"user_one" json:"user_one"`
	UserTwo      models.User `bson:"user_two" json:"user_two"`
	CreationDate time.Time   `bson:"creation_date" json:"creation_date"`
}

type ConnectionsQuery []*ConnectionQuery

func (c *ConnectionInput) New() *ConnectionInput {
	c.ID = uuid.New()
	c.CreationDate = time.Now()
	return c
}
