package frontweather

import "github.com/gofiber/fiber/v2"

func Weather(message string) {
	app := fiber.New()

	app.Get("/weather", func(c *fiber.Ctx) error {
		return c.SendString(message)
	})

	app.Listen(":3000")
}
