package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hemanth5603/RuleEngineBackend/config"
	"github.com/hemanth5603/RuleEngineBackend/models"
	"github.com/hemanth5603/RuleEngineBackend/utils"
)

// Request structure for combining rules

func CombineRuleAPI(c *fiber.Ctx) error {
	var req models.CombineRulesRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request payload")
	}

	if len(req.RootNodeIds) != 2 {
		return c.Status(http.StatusBadRequest).SendString("Exactly two root node IDs must be provided")
	}

	leftExpr, err := buildExpression(req.RootNodeIds[0])
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error building left rule expression")
	}

	rightExpr, err := buildExpression(req.RootNodeIds[1])
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error building right rule expression")
	}

	rule1, _ := utils.CreateRule(leftExpr)

	rule2, _ := utils.CreateRule(rightExpr)

	combinedRule, _ := utils.CombineRules([]*models.Node{rule1, rule2}, req.Operator)

	id, err := utils.DBSaveRule(config.POSTGRES_DB, combinedRule)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err})
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "id": id})

}

func buildExpression(nodeID int) (string, error) {
	node, err := utils.DBfetchNodeByID(config.POSTGRES_DB, nodeID)
	if err != nil {
		return "", err
	}

	// If it's an operand (leaf node), return the condition
	if node.NodeType == "operand" {
		return fmt.Sprintf("%s %s %s", node.Attribute, node.Operator, node.Value), nil
	}

	// If it's an operator, recursively build the expression for left and right children
	var leftExpr, rightExpr string
	if node.LeftChild != nil {
		leftExpr, err = buildExpression(*node.LeftChild)
		if err != nil {
			return "", err
		}
	}

	if node.RightChild != nil {
		rightExpr, err = buildExpression(*node.RightChild)
		if err != nil {
			return "", err
		}
	}

	// Combine left and right expressions with the operator (AND/OR)
	return fmt.Sprintf("(%s %s %s)", leftExpr, node.NodeType, rightExpr), nil
}
