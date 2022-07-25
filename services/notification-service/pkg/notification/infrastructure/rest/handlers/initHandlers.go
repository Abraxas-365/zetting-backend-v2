package handlers

import (
	"notifications/pkg/notification/application"

	"github.com/gofiber/fiber/v2"
)

type NotificationHandler interface {
	Create(c *fiber.Ctx) error
	FindByUserID(c *fiber.Ctx) error
	ChangeNotificationStatus(c *fiber.Ctx) error
}

type handler struct {
	app application.NotificationApp
}

func NewNoticationHandler(app application.NotificationApp) NotificationHandler {
	return &handler{
		app,
	}

}
