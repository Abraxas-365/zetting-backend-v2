package ports

import (
	"notifications/pkg/notification/domain/models"

	"github.com/google/uuid"
)

type NotificationRepo interface {
	Create(new models.NotificationInput) (bool, error)
	FindByUserID(id uuid.UUID, page int) (models.NotificationsQuery, bool, error) //Get all notifications for a user
	ChangeNotificationStatus(notificationId uuid.UUID, isSeen bool) error
}
