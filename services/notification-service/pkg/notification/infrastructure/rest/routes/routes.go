package routes

import (
	"notifications/pkg/notification/infrastructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func NotificationsRoutes(app *fiber.App, handler handlers.NotificationHandler) {
	users := app.Group("/api/notifications")
	/*Login user*/
	users.Post("/create", handler.Create)
	/*Register user*/
	users.Get("/find/userid=:id", handler.FindByUserID)
	/*Get user by id*/
	users.Put("/update", handler.ChangeNotificationStatus)
}
