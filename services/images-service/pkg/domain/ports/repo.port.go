package ports

import (
	"service-images/pkg/domain/events"
	"service-images/pkg/domain/models"
)

type Repository interface {
	PostImage(image models.Image) (events.Event, error)
	GetImagesFilter(images models.ImageFilter, page int) (models.ImageList, error)

	// SaveAlbum(album models.Album) error
	// AddImagesToAlbum(Album uuid.UUID, user uuid.UUID, images models.ImageList) error
}
