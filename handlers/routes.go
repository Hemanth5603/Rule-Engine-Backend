package handlers

import "github.com/gofiber/fiber/v2"

func AppRoutes(incommingRoutes *fiber.App) {
	incommingRoutes.Post("/api/create-rule", CreateRuleAPI)
}
