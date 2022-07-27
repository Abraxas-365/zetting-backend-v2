package application

import (
	"connection-service/pkg/connection/domain/models"
	"connection-service/pkg/connection/domain/ports"

	"github.com/google/uuid"
)

type Application interface {
	Create(new models.ConnectionInput) (bool, error)
	FindByUserID(userId uuid.UUID) (models.ConnectionsQuery, bool, error)
	DeleteConection(userId uuid.UUID) error
}

type application struct {
	connectionRepo ports.ConnectionRepo
}

func NewApp(connectionRepo ports.ConnectionRepo) Application {
	return &application{
		connectionRepo,
	}
}
