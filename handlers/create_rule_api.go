package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hemanth5603/RuleEngineBackend/config"
	"github.com/hemanth5603/RuleEngineBackend/models"
	"github.com/hemanth5603/RuleEngineBackend/utils"
)

func CreateRuleAPI(ctx *fiber.Ctx) error {
	var payload models.CreateRuleRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	root, err := utils.CreateRule(payload.Rule)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err})
	}

	id, err := utils.DBSaveRule(config.POSTGRES_DB, root)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": false, "error": err})
	}

	return ctx.Status(fiber.StatusAccepted).
		JSON(fiber.Map{"status": true, "id": id})

}
