package application

import (
	"mime/multipart"
	"service-images/pkg/domain/models"
)

func (app *application) PostImage(file multipart.File, image models.Image) error {
	uploaded, err := app.bucket.Upload(file, image.ID.String())
	if err != nil {
		return err
	}
	image.EntityTag = *uploaded.ETag
	image.UrlPath = *&uploaded.Location

	_, err = app.repo.PostImage(image)
	if err != nil {
		return err
	}
	return nil
}
