package models

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type Image struct {
	ID        uuid.UUID     `bson:"_id,omitempty" json:"id"`
	EntityTag string        `bson:"entity_tag,omitempty" json:"entity_tag"`
	Name      string        `bson:"name" json:"name"`
	Owner     uuid.UUID     `bson:"owner" json:"owner"`
	Tag       string        `bson:"tag" json:"tag"`
	Metadata  ImageMetadata `bson:"metadata" json:"metadata"`
	UrlPath   string        `bson:"url_path" json:"url_path"`
}

type ImageMetadata struct {
	Format   Format `bson:"format" json:"format"`
	Size     int    `bson:"size" json:"size"`
	Location string `bson:"location" json:"location"`
	//TODO more metadata metrics
}

type ImageList []*Image

func NewImage(file *multipart.FileHeader,
	owner uuid.UUID,
	tag string,
) Image {
	i := &Image{}

	i.ID = uuid.New()
	i.Tag = tag
	i.Name = file.Filename
	i.Owner = owner
	i.Metadata.Size = int(file.Size)
	i.Metadata.Format = i.Metadata.Format.GetFormat(i.Name)
	return *i
}
