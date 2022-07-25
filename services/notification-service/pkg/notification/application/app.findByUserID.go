package application

import (
	"notifications/pkg/notification/domain/models"

	"github.com/google/uuid"
)

func (a *notificationApp) FindByUserID(id uuid.UUID) ([]models.NotificationsQuery, error) {
	return a.FindByUserID(id)
}
