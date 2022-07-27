package models

import "github.com/google/uuid"

type User struct {
	ID     uuid.UUID `bson:"_id,omitempty" json:"id"`
	Name   Name      `bson:"name" json:"name"`
	Avatar string    `bson:"avatar" json:"avatar"`
}

type Name struct {
	FirstName  string `bson:"first_name" json:"first_name,omitempty"`
	MiddleName string `bson:"middle_name" json:"middle_name,omitempty"`
	LastName   string `bson:"last_name" json:"last_name"`
}
