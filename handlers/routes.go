package handlers

import "github.com/gofiber/fiber/v2"

func AppRoutes(incommingRoutes *fiber.App) {
	incommingRoutes.Post("/api/create-rule", CreateRuleAPI)
	incommingRoutes.Get("/api/get-all-nodes", FetchRules)

	incommingRoutes.Post("/api/combine-rules", CombineRuleAPI)
	incommingRoutes.Post("/api/evaluate-rules", EvaluateRuleAPI)
}
