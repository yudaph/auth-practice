package routes

import (
	"task/package/entities"
	"task/package/users"
	"task/zapLog"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service users.Service) {
	app.Get("/", getUsers(service))
	app.Post("/", createUser(service))
	app.Get("/:id", getUser(service))
	app.Patch("/:id", updateUser(service))
	app.Delete("/:id", deleteUser(service))
}

func getUsers(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.GetUsers(&map[string]interface{}{
			"status": "on",
		})
		if err != nil {
			return c.Status(err.Code).JSON(fiber.Map{
				"message" : err.GetMessage().Message,
			})
		}

		return c.Status(200).JSON(&fetched)
	}
}

func getUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.GetUser(&map[string]interface{}{
			"id": c.Params("id"),
			"status": "on",
		})
		if err != nil {
			return c.Status(err.Code).JSON(fiber.Map{
				"message" : err.GetMessage().Message,
			})
		}

		return c.Status(200).JSON(&fetched)
	}
}

func createUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(entities.UserRegister)
		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser " + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}

		user, err := service.InsertUser(body)
		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(&user)
	}
}

func updateUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(map[string]interface{})
		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser " + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}

		id := c.Params("id")
		err := service.UpdateUser(body, &id)
		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "update success",
		})
	}
}

func deleteUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := service.RemoveUser(&id); err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "delete success",
		})
	}
}