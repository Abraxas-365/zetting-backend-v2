package handlers

import (
	"fmt"
	"service-users/pkg/user/domain/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) Create(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"timestamp": time.Now(),
			"error":     fiber.ErrBadRequest.Message,
			"message":   err.Error(),
		})
	}
	fmt.Println("newUser", newUser)
	if check, err := h.userApp.Create(*newUser.New()); !check {
		return c.Status(202).JSON(fiber.Map{
			"timestamp": time.Now(),
			"error":     fiber.ErrBadRequest.Message,
		})

	} else if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(201)
}
