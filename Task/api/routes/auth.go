package routes

import (
	"task/package/auth"
	"task/package/entities"
	"task/zapLog"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service auth.Service) {
	app.Post("/", login(service))
	app.Post("/change-password", Protect(service), changePassword(service))
}

func login(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(entities.Login)
		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser " + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Internal Server Error body parser")
		}

		res, err := service.Login(body)
		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(&res)
	}
}

func changePassword(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(entities.ChangePassword)
		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser " + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}

		body.ID = string(c.Request().Header.Peek("ID"))

		if err:=service.ChangePassword(body); err!=nil {
			return fiber.NewError(err.Code, err.Message)
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "Password succesfuly change",
		})
	}
}

func Protect(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Request().Header.Peek("Authorization")
		if authorization == nil {
			zapLog.Error("Not authorized")
			return fiber.NewError(fiber.StatusUnauthorized, "Please login")
		}
		authstring := string(authorization)
		
		res, err := service.Auth(&authstring)
		if err!=nil{
			return fiber.NewError(err.Code, err.Message)
		}

		c.Request().Header.Add("ID", *res)

		return c.Next()
	}
}