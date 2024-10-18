package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hemanth5603/RuleEngineBackend/config"
	"github.com/hemanth5603/RuleEngineBackend/utils"
)

func FetchRules(ctx *fiber.Ctx) error {

	expressions, err := utils.BuildExpressionsForAllNodes(config.POSTGRES_DB)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error building expressions")
	}

	// Return the expressions as JSON
	return ctx.JSON(fiber.Map{
		"expressions": expressions,
	})
}
