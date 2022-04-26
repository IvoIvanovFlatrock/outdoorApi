package main

import (
	"log"
	"outdoorsy_api/database"
	"outdoorsy_api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file. Using OS environment")
	}

	app := fiber.New()

	database.Connect()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
