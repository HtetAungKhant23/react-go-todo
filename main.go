package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hi gophers!")
	app := fiber.New()

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "server still alive."})
	})

	log.Fatal(app.Listen(":4000"))
}
