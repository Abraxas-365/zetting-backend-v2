package handler

import (
	"service-images/pkg/domain/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetImagesByFilter(c *fiber.Ctx) error {
	body := new(models.ImageFilter)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"timeStamp": time.Now(),
			"error":     true,
			"msg":       err.Error(),
		})
	}
	listIamges, err := h.app.GetImagesFilter(*body, 1)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"timeStamp": time.Now(),
			"error":     true,
			"msg":       err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(listIamges)
}
