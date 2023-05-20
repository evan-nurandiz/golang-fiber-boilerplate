package validation

import (
	"github.com/evan_nurandiz/go_fiber_boilerplate/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

type IError struct {
	Field string
	Tag   string
	Value string
}

type Register struct {
	Name            string `json:"name" validate:"required,min=1,max=30"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=3,max=24"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=24"`
}

func ValidateRegister(c *fiber.Ctx) error {
	body := new(Register)
	c.BodyParser(&body)

	err := ValidateRequest(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if body.ConfirmPassword != body.Password {
		response := helpers.BuildErrorResponse(
			fiber.StatusBadRequest, "Failed to process request", "password not match")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Next()
}

func ValidateLogin(c *fiber.Ctx) error {
	body := new(Login)
	c.BodyParser(&body)

	err := ValidateRequest(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Next()
}

func ValidateRequest(body interface{}) interface{} {
	var errors []*IError
	err := Validator.Struct(body)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}

		return errors
	}

	return nil
}
