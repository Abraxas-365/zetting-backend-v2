package application

import (
	"notifications/pkg/notification/domain/models"
	"notifications/pkg/notification/domain/ports"

	"github.com/google/uuid"
)

type NotificationApp interface {
	Create(new models.NotificationInput) (bool, error)
	FindByUserID(id uuid.UUID, page int) (models.NotificationsQuery, error) //Get all notifications for a user
	ChangeNotificationStatus(notificationId uuid.UUID, isSeen bool) error
}

type notificationApp struct {
	notificationRepo ports.NotificationRepo
}

func NewNotificationApp(notificationRepo ports.NotificationRepo) NotificationApp {
	return &notificationApp{
		notificationRepo,
	}
}
