package handler

import (
	"service-images/internal/auth"
	"service-images/pkg/domain/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) UploadImage(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"timeStamp": time.Now(),
			"error":     true,
			"msg":       err.Error(),
		})
	}
	// Get Buffer from file
	buffer, err := file.Open()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"timeStamp": time.Now(),
			"error":     true,
			"msg":       err.Error(),
		})
	}
	defer buffer.Close()

	// Upload the zip file with PutObject
	tag := c.FormValue("tag", "no-tag")
	if err := h.app.PostImage(buffer, models.NewImage(file, userTokenData.ID, tag)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"timeStamp": time.Now(),
			"error":     true,
			"msg":       err.Error(),
		})
	}

	return c.SendStatus(201)
}
