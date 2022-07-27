package repository

import (
	"connection-service/pkg/connection/domain/models"

	"github.com/google/uuid"
)

func (r *repository) FindByUserID(userId uuid.UUID) (models.ConnectionsQuery, bool, error) {

	return models.ConnectionsQuery{}, true, nil
}
