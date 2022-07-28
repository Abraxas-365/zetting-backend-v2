package application

import (
	"notifications/pkg/notification/domain/models"

	"github.com/google/uuid"
)

func (a *notificationApp) FindByUserID(id uuid.UUID, page int) (models.NotificationsQuery, error) {
	notifications, work, err := a.notificationRepo.FindByUserID(id, page)
	if !work {
		return nil, err
	}
	return notifications, nil

}
