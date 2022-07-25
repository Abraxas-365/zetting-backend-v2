package application

import "github.com/google/uuid"

func (a *notificationApp) ChangeNotificationStatus(notificationId uuid.UUID, isSeen bool) error {
	return a.notificationRepo.ChangeNotificationStatus(notificationId, isSeen)
}
