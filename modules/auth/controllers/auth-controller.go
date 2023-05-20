package controllers

import (
	"fmt"

	"github.com/evan_nurandiz/go_fiber_boilerplate/helpers"
	"github.com/evan_nurandiz/go_fiber_boilerplate/modules/auth/handlers"
	"github.com/evan_nurandiz/go_fiber_boilerplate/modules/auth/model"
	"github.com/evan_nurandiz/go_fiber_boilerplate/services"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	user := new(model.User)
	c.BodyParser(user)

	existUser, err := handlers.GetUserDataByEmail(user.Email)

	if err != "" {
		errMessage := helpers.MappingError(err, "auth")
		response := helpers.BuildErrorResponse(
			fiber.StatusBadRequest, "Failed to process request", errMessage)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	checkPassword := helpers.VerifyPassword(existUser.Password, []byte(user.Password))

	type LoginResponse struct {
		User_id       int    `json:"user_id"`
		Email         string `json:"email"`
		Access_token  string `json:"token"`
		Refresh_token string `json:"refresh_token"`
	}

	fmt.Println(existUser)

	if checkPassword {
		err, accessToken, refreshToken := services.GenerateToken(services.JWTpayload{
			Email:   existUser.Email,
			User_id: existUser.User_id,
			Name:    existUser.Name,
		})

		if err != "" {
			fmt.Println(err)
			panic(err)
		}

		response := &LoginResponse{
			User_id:       existUser.User_id,
			Email:         existUser.Email,
			Access_token:  accessToken,
			Refresh_token: refreshToken,
		}

		return c.Status(fiber.StatusAccepted).JSON(helpers.BuildResponse(fiber.StatusAccepted, "success login", response))
	}

	response := helpers.BuildErrorResponse(
		fiber.StatusBadRequest, "Failed to process request", "password not match")
	return c.Status(fiber.StatusBadRequest).JSON(response)
}

func Register(c *fiber.Ctx) error {
	user := new(model.User)
	c.BodyParser(user)
	createdUser, errCreatedUser := handlers.RegisterUser(*user)

	if errCreatedUser != nil {
		errMessage := helpers.MappingError(errCreatedUser.Error(), "auth")
		response := helpers.BuildErrorResponse(
			fiber.StatusBadRequest, "Failed to process request", errMessage)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helpers.BuildResponse(fiber.StatusAccepted, "success register", createdUser)

	return c.Status(fiber.StatusAccepted).JSON(response)
}
