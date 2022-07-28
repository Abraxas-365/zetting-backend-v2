package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handler) FindByUserID(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"timestamp": time.Now(),
			"error":     fiber.ErrBadRequest.Message,
		})
	}
	//TODO make page variable
	notifications, err := h.app.FindByUserID(userId, 1)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"timestamp": time.Now(),
			"error":     fiber.ErrBadRequest.Message,
		})
	}
	return c.Status(fiber.StatusOK).JSON(notifications)

}
