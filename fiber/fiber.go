package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/a", func(c *fiber.Ctx) error {
		filepath := c.Params("filepath")
		return c.SendFile(filepath)
	})
	app.Get("/b", func(c *fiber.Ctx) error {
		filepath := c.Params("filepath")
		header, _ := c.FormFile("f")
		_ = c.SaveFile(header, filepath)
		c.Attachment(filepath)
		return c.SendFile(filepath)
	})
	app.Listen(":3000")
}
