package application

import (
	"mime/multipart"
	"service-images/pkg/domain/models"
	"service-images/pkg/domain/ports"
)

type Application interface {
	PostImage(file multipart.File, image models.Image) error
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
