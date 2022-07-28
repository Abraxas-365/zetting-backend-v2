package application

import (
	"fmt"
	"notifications/pkg/notification/domain/models"
)

func (a *notificationApp) Create(new models.NotificationInput) (bool, error) {
	fmt.Println("create application")
	return a.notificationRepo.Create(*new.New())
}
