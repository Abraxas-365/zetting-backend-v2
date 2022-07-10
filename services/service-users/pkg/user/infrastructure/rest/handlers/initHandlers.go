package handlers

import (
	"service-users/pkg/user/application"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	GetPublicUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type handler struct {
	userApp application.Application
}

func NewHandler(app application.Application) Handler {
	return &handler{
		app,
	}
}
