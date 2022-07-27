package application

import "connection-service/pkg/connection/domain/models"

func (a *application) Create(new models.ConnectionInput) (bool, error) {
	return a.connectionRepo.Create(new)
}
