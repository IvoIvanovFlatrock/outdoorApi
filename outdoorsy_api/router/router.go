package router

import (
	rentalHandler "outdoorsy_api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	rental := router.Group("/rentals")

	rental.Get("/:price_min?", func(c *fiber.Ctx) error {
		return rentalHandler.GetRentalsWithParams(c)
	})

	rental.Get("/", func(c *fiber.Ctx) error {
		return rentalHandler.GetRentals(c)
	})

}
