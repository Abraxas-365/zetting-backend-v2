package application

import (
	"connection-service/pkg/connection/domain/models"

	"github.com/google/uuid"
)

func (a *application) FindByUserID(userId uuid.UUID) (models.ConnectionsQuery, bool, error) {
	return a.connectionRepo.FindByUserID(userId)
}
