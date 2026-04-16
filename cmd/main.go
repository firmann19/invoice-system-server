package main

import (
	"log"
	"os"

	"fleetify-test/config"
	"fleetify-test/models"
	"fleetify-test/database"
	"fleetify-test/src/handlers"
	"fleetify-test/src/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app.Post("/api/login", handlers.Login)

	app.Post("/api/invoices", middlewares.JWTMiddleware, handlers.CreateInvoice)

	app.Get("/api/items", handlers.GetItemByCode)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running 🚀")
	})

	port := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + port))
}