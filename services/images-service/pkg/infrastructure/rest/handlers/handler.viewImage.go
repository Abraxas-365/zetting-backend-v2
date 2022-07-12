package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) ViewImage(c *fiber.Ctx) error {
	imageId := c.Params("id")
	buff, err := h.app.ViewImage(imageId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"timeStamp": time.Now(),
			"error":     true,
			"msg":       err.Error(),
		})
	}
	return c.Send(buff.Bytes())
}
