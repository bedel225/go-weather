package frontweather

import (
	"github.com/gofiber/fiber/v2"
)

func Index() {
	app := fiber.New()
	message := "bienvenue sur votre site de la meteo"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(message)
	})

	app.Listen(":3000")
}
