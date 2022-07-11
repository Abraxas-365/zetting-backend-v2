package models

import "github.com/google/uuid"

type Album struct {
	ID     uuid.UUID `bson:"_id,omitempty" json:"id"`
	Owner  uuid.UUID `bson:"owner" json:"owner"`
	Name   string    `bson:"name" json:"name"`
	Images ImageList `bson:"images" json:"images"`
}

type AlbumList []*Album
