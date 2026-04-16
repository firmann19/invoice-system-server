package handlers

import (
	"github.com/gofiber/fiber/v2"
	"fleetify-test/config"
	"fleetify-test/src/services"
)

func GetItemByCode(c *fiber.Ctx) error {
	code := c.Query("code")

	if code == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "code is required",
		})
	}

	item, err := services.FindItemByCode(config.DB, code)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":    item.ID,
		"code":  item.Code,
		"name":  item.Name,
		"price": item.Price,
	})
}