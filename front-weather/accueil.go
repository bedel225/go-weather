package frontweather

import "github.com/gofiber/fiber/v2"

func Index(message string) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(message)
	})

	app.Listen(":3000")
}
