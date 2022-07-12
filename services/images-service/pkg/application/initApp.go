package application

import (
	"bytes"
	"mime/multipart"
	"service-images/pkg/domain/models"
	"service-images/pkg/domain/ports"
)

type Application interface {
	PostImage(file multipart.File, image models.Image) error
	GetImage(key string) (*bytes.Buffer, error)
	GetImagesByTag(tag string) (models.ImageList, error)
}

type application struct {
	repo   ports.Repository
	bucket ports.BucketRepo
}

func NewApplication(repo ports.Repository, bucket ports.BucketRepo) Application {
	return &application{
		repo,
		bucket,
	}
}
