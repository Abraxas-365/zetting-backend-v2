package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `bson:"_id,omitempty" json:"id"`
	Name     string    `bson:"name" json:"name"`
	LastName string    `bson:"last_name" json:"last_name"`
	Avatar   string    `bson:"avatar" json:"avatar"`
}
