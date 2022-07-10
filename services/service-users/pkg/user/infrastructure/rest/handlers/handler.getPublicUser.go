package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handler) GetPublicUser(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	user, check, err := h.userApp.GetPublicUser(userId)
	if !check {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
