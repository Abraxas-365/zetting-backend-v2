package application

import (
	"mime/multipart"
	"service-images/pkg/domain/models"
)

func (app *application) PostImage(file multipart.File, image models.Image) error {
	//image.ID.String()+"."+string(image.Metadata.Format) this put the format to the file
	uploaded, err := app.bucket.Upload(file, image.ID.String()+"."+string(image.Metadata.Format))
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
