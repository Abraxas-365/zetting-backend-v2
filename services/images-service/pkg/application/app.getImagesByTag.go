package application

import "service-images/pkg/domain/models"

func (app *application) GetImagesByTag(tag string) (models.ImageList, error) {
	return app.repo.GetImagesByTag(tag)

}
