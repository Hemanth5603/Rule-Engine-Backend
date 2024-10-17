package main

import (
	"fmt"

	"github.com/hemanth5603/RuleEngineBackend/models"
	"github.com/hemanth5603/RuleEngineBackend/utils"
)

func main() {
	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	// config.InitializePostgresSQL()

	rule1, _ := utils.CreateRule("age > 30 AND department = 'Sales'")

	rule2, _ := utils.CreateRule("experience > 5 OR salary > 50000")

	combinedRule, _ := utils.CombineRules([]*models.Node{rule1, rule2}, "AND")

	user := models.UserModel{
		Age:        52,
		Department: "Sales",
		Salary:     45000,
		Experience: 4,
	}

	result := utils.EvaluateRule(combinedRule, user)

	fmt.Print(result)

	//app.Listen(":8081")
}
