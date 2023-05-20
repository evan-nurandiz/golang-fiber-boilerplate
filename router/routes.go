package router

import (
	"github.com/evan_nurandiz/go_fiber_boilerplate/modules/auth/routes"
	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	authRouter := api.Group("/auth")

	//auth
	routes.AuthRoutes(authRouter)
}
