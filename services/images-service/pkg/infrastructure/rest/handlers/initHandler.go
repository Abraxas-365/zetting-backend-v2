package handler

import (
	"service-images/pkg/application"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	UploadImage(c *fiber.Ctx) error
}

type handler struct {
	app application.Application
}

func NewHandler(app application.Application) Handler {
	return &handler{
		app,
	}
}
