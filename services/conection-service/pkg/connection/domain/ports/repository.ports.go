package ports

import (
	"connection-service/pkg/connection/domain/models"

	"github.com/google/uuid"
)

type ConnectionRepo interface {
	Create(new models.ConnectionInput) (bool, error)
	FindByUserID(userId uuid.UUID) (models.ConnectionsQuery, bool, error)
	DeleteConection(userId uuid.UUID) error
}
