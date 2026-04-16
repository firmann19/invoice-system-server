package main

import (
	"log"
	"os"

	"fleetify-test/config"
	"fleetify-test/models"
	"fleetify-test/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.Item{},
		&models.Invoice{},
		&models.InvoiceDetail{},
	)

	database.SeedItems(config.DB)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running 🚀")
	})

	port := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + port))
}