package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/ross714/hcaptcha"
)

func main() {
	client := hcaptcha.New("secret_key", "site_key")

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		token := c.FormValue("h-captcha-response")
		res := client.Verify(token)

		if res {
			return c.SendString("success")
		} else {
			return c.SendString("failed")
		}
	})

	app.Listen(":3000")
}
