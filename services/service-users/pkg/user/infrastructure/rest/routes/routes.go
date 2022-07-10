package routes

import (
	"service-users/pkg/user/infrastructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler handlers.Handler) {
	users := app.Group("/api/users")
	/*Login user*/
	users.Post("/login", handler.Login)
	/*Register user*/
	users.Post("/register", handler.Create)
	/*Get user by id*/
	users.Get("/id=:id", handler.GetPublicUser)
}
