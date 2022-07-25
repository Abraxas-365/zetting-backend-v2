package handlers

import (
	"notifications/pkg/notification/domain/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) Create(c *fiber.Ctx) error {

	newNotification := new(models.NotificationInput)
	if err := c.BodyParser(newNotification); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"timestamp": time.Now(),
			"error":     fiber.ErrBadRequest.Message,
			"message":   err.Error(),
		})
	}
	if check, err := h.app.Create(*newNotification.New()); !check {
		return c.Status(202).JSON(fiber.Map{
			"timestamp": time.Now(),
			"error":     fiber.ErrBadRequest.Message,
		})

	} else if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(201)

}
