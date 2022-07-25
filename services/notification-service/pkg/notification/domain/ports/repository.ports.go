package ports

import (
	"notifications/pkg/notification/domain/models"

	"github.com/google/uuid"
)

type NotificationRepo interface {
	Create(new models.NotificationInput) (bool, error)
	FindByUserID(id uuid.UUID) ([]models.NotificationsQuery, error) //Get all notifications for a user
	ChangeNotificationStatus(notificationId uuid.UUID, isSeen bool) error
}
