package routes

import (
	"service-images/internal/auth"
	handler "service-images/pkg/infrastructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func ServeRoute(app *fiber.App, handler handler.Handler) {
	/*SERVE*/
	static := app.Group("/api/static")
	static.Post("/upload", auth.JWTProtected(), handler.UploadImage)
	static.Get("/id=:id", handler.GetImage)
	static.Get("/tag=:tag", handler.GetImagesByTag)
}
