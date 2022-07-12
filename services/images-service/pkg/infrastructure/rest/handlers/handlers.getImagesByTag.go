package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetImagesByTag(c *fiber.Ctx) error {
	tag := c.Params("tag", "")
	listIamges, err := h.app.GetImagesByTag(tag)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"timeStamp": time.Now(),
			"error":     true,
			"msg":       err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(listIamges)
}
