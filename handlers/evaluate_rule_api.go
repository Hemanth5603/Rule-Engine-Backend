package handlers

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hemanth5603/RuleEngineBackend/config"
	"github.com/hemanth5603/RuleEngineBackend/models"
	"github.com/hemanth5603/RuleEngineBackend/utils"
	"github.com/hemanth5603/RuleEngineBackend/workers"
)

func CreateNode(condition string) *models.Node {

	parts := strings.Split(condition, " ")
	if len(parts) == 3 {
		return &models.Node{
			NodeType: "operand",
			Value: &models.Condition{
				Attribute: parts[0],
				Operator:  parts[1],
				Value:     parseValue(parts[2]),
			},
		}
	}
	return nil
}

func parseValue(value string) interface{} {
	if intVal, err := strconv.Atoi(value); err == nil {
		return intVal
	}
	return value
}

func CombineNodes(left *models.Node, right *models.Node, operator string) *models.Node {
	return &models.Node{
		NodeType: operator,
		Left:     left,
		Right:    right,
	}
}

func RebuildRuleTree(rules []string, operators []string) *models.Node {
	if len(rules) == 0 || len(operators) == 0 {
		return nil
	}

	// Start with the first rule as the left node
	root := CreateNode(rules[0])

	// Loop through the remaining rules and operators to combine them into a tree
	for i := 1; i < len(rules); i++ {
		operator := operators[i-1]
		right := CreateNode(rules[i])
		root = CombineNodes(root, right, operator) // Combine the current root with the next rule
	}

	return root
}

func BuildExpressionWithSubRules(db *sql.DB, rootID int, subRules *[]string) (string, error) {
	// Fetch the root node by its ID
	rootNode, err := utils.DBfetchNodeByID(db, rootID)
	if err != nil {
		return "", err
	}

	if rootNode.NodeType == "operand" {
		return fmt.Sprintf("%s %s %s", rootNode.Attribute, rootNode.Operator, rootNode.Value), nil
	}

	var leftExpr, rightExpr string

	if rootNode.LeftChild != nil {
		leftExpr, err = BuildExpressionWithSubRules(db, *rootNode.LeftChild, subRules)
		if err != nil {
			return "", err
		}
	}

	if rootNode.RightChild != nil {
		rightExpr, err = BuildExpressionWithSubRules(db, *rootNode.RightChild, subRules)
		if err != nil {
			return "", err
		}
	}

	combinedExpr := fmt.Sprintf("(%s %s %s)", leftExpr, rootNode.NodeType, rightExpr)

	// Add this combined expression as a "sub-rule" if it's a valid sub-expression (AND/OR node)
	*subRules = append(*subRules, combinedExpr)

	return combinedExpr, nil
}

func FetchAllSubRules(db *sql.DB, rootID int) ([]string, string, error) {
	var subRules []string
	combinedExpr, err := BuildExpressionWithSubRules(db, rootID, &subRules)
	if err != nil {
		return nil, "", err
	}
	return subRules, combinedExpr, nil
}

func BuildExpressionWithRules(db *sql.DB, rootID int, rules *[]string) (string, error) {
	rootNode, err := utils.DBfetchNodeByID(db, rootID)
	if err != nil {
		return "", err
	}

	if rootNode.NodeType == "operand" {
		rule := fmt.Sprintf("%s %s %s", rootNode.Attribute, rootNode.Operator, rootNode.Value)
		*rules = append(*rules, rule) // Add the rule to the array
		return rule, nil
	}

	var leftExpr, rightExpr string

	if rootNode.LeftChild != nil {
		leftExpr, err = BuildExpressionWithRules(db, *rootNode.LeftChild, rules)
		if err != nil {
			return "", err
		}
	}

	if rootNode.RightChild != nil {
		rightExpr, err = BuildExpressionWithRules(db, *rootNode.RightChild, rules)
		if err != nil {
			return "", err
		}
	}

	combinedExpr := fmt.Sprintf("(%s %s %s)", leftExpr, rootNode.NodeType, rightExpr)

	return combinedExpr, nil
}

func FetchAllRules(db *sql.DB, rootID int) ([]string, string, error) {
	var rules []string
	combinedExpr, err := BuildExpressionWithRules(db, rootID, &rules)
	if err != nil {
		return nil, "", err
	}
	return rules, combinedExpr, nil
}

func FetchLastRecordID(db *sql.DB) (int, error) {
	var lastID int
	query := "SELECT id FROM ast_nodes ORDER BY id DESC LIMIT 1"

	err := db.QueryRow(query).Scan(&lastID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no records found in the table")
		}
		return 0, err
	}
	return lastID, nil
}

func EvaluateRuleAPI(c *fiber.Ctx) error {
	var payload models.EvaluateRuleRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	id, err := FetchLastRecordID(config.POSTGRES_DB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": false, "err": err})
	}

	rules, _, err := FetchAllRules(config.POSTGRES_DB, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": false, "error": err})
	}

	_, combinedRule, err := FetchAllSubRules(config.POSTGRES_DB, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": false, "error": err})
	}

	for i, rule := range rules {
		// Remove all instances of '(' and ')'
		rules[i] = strings.ReplaceAll(rule, "(", "")
		rules[i] = strings.ReplaceAll(rules[i], ")", "")
	}

	fmt.Println(rules)
	fmt.Println(len(rules))

	fmt.Println(combinedRule)
	combinedRule = strings.ReplaceAll(combinedRule, "(", "")
	combinedRule = strings.ReplaceAll(combinedRule, ")", "")
	fmt.Println(combinedRule)
	operators := workers.CustomExtract(combinedRule)
	// for i := 0; i < len(oper); i++ {
	// 	if oper[i] == "AND" || oper[i] == "OR" {
	// 		operators = append(operators, oper[i])
	// 	}
	// }

	// for i, j := 0, len(operators)-1; i < j; i, j = i+1, j-1 {
	// 	operators[i], operators[j] = operators[j], operators[i]
	// }

	// for i := 0; i < len(rules); i++ {
	// 	op := strings.Fields(rules[i])
	// 	for j := 0; j < len(op); j++ {
	// 		if op[j] == "AND" || op[j] == "OR" {
	// 			operators = append(operators, op[j])
	// 		}
	// 	}
	// }

	fmt.Println(operators)

	rootNode := RebuildRuleTree(rules, operators)

	user := models.UserModel{
		Age:        payload.Age,
		Salary:     payload.Salary,
		Experience: payload.Experience,
		Department: payload.Department,
	}

	result := utils.EvaluateRule(rootNode, user)

	if result {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{"evaluation": result, "message": "User Valid"})
	}
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{"evaluation": result, "message": "User InValid"})

}
