package application

import "notifications/pkg/notification/domain/models"

func (a *notificationApp) Create(new models.NotificationInput) (bool, error) {
	return a.notificationRepo.Create(*new.New())
}
