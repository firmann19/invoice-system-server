package handlers

import (
	"github.com/gofiber/fiber/v2"
	"fleetify-test/src/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	for _, user := range auth.Users {
		if user.Username == req.Username && user.Password == req.Password {
			token, err := auth.GenerateToken(user)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": "Failed generate token",
				})
			}

			return c.JSON(fiber.Map{
				"token": token,
			})
		}
	}

	return c.Status(401).JSON(fiber.Map{
		"message": "Invalid credentials",
	})
}