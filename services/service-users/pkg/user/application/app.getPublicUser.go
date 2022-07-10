package application

import (
	"service-users/pkg/user/domain/models"

	"github.com/google/uuid"
)

func (app *application) GetPublicUser(uuid uuid.UUID) (models.UserPublic, bool, error) {
	return app.userRepo.GetPublicUser(uuid)
}
