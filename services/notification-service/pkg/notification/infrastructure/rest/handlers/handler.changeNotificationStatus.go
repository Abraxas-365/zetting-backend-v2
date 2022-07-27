package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handler) ChangeNotificationStatus(c *fiber.Ctx) error {
	body := struct {
		NotificationId uuid.UUID `json:"notification_id"`
		Status         bool      `json:"status"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"timestamp": time.Now(),
			"error":     fiber.ErrBadRequest.Message,
		})
	}
	if err := h.app.ChangeNotificationStatus(body.NotificationId, body.Status); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"timestamp": time.Now(),
			"error":     fiber.ErrBadRequest.Message,
		})
	}
	return c.SendStatus(fiber.StatusOK)

}
