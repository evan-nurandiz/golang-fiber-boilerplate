package routes

import (
	"github.com/evan_nurandiz/go_fiber_boilerplate/modules/auth/controllers"
	"github.com/evan_nurandiz/go_fiber_boilerplate/modules/auth/validation"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	api.Post("/login", validation.ValidateLogin, controllers.Login)
	api.Post("/register", validation.ValidateRegister, controllers.Register)
}
