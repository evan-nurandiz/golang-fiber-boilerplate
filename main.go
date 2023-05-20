package main

import (
	"fmt"

	"github.com/evan_nurandiz/go_fiber_boilerplate/config"
	"github.com/evan_nurandiz/go_fiber_boilerplate/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	router.Routes(api)

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello wordldss")
	})

	app.Listen(fmt.Sprintf(":%s",
		config.GetConfig("PORT"),
	))
}
