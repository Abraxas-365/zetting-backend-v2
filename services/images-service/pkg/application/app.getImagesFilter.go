package application

import "service-images/pkg/domain/models"

func (app *application) GetImagesFilter(images models.ImageFilter, page int) (models.ImageList, error) {
	return app.repo.GetImagesFilter(images, page)

}
