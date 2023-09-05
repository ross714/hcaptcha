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
		return c.Render("index", fiber.Map{
			"Sitekey": client.Sitekey,
		})
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		token := c.FormValue("h-captcha-response")
		res := client.Verify(token)

		if res {
			return c.SendString("Success")
		} else {
			return c.SendString("Failed")
		}
	})

	app.Listen(":5000")
}
