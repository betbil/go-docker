package handlers

import (
	"github.com/betbil/go-docker/database"
	"github.com/betbil/go-docker/models"
	"github.com/gofiber/fiber/v2"
)

func HomeHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World hii4!")
}

func CreateFactHandler(ctx *fiber.Ctx) error {
	fact := new(models.Fact)
	err := ctx.BodyParser(fact)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	database.DB.Db.Create(&fact)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "success", "data": fact})
}

func ListFactHandler(ctx *fiber.Ctx) error {
	var facts []models.Fact
	database.DB.Db.Find(&facts)
	if len(facts) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No facts found", "data": nil})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "success", "data": facts})
}
