package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hemanth5603/RuleEngineBackend/config"
	"github.com/hemanth5603/RuleEngineBackend/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	handlers.AppRoutes(app)
	config.InitializePostgresSQL()

	// rule1, _ := utils.CreateRule("age > 30 AND department = 'Sales'")

	// rule2, _ := utils.CreateRule("experience > 5 OR salary > 50000")

	// combinedRule, _ := utils.CombineRules([]*models.Node{rule1, rule2}, "AND")

	// user := models.UserModel{
	// 	Age:        52,
	// 	Department: "Sales",
	// 	Salary:     45000,
	// 	Experience: 4,
	// }

	// result := utils.EvaluateRule(combinedRule, user)

	// fmt.Print(result)

	app.Listen(":8081")
}
