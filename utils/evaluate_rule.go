package utils

import (
	"fmt"

	"github.com/hemanth5603/RuleEngineBackend/models"
)

// EvaluateRule traverses the AST and evaluates the rule against the provided user data
func EvaluateRule(node *models.Node, userData models.UserModel) bool {

	if node == nil {
		return false
	}

	// If it's an AND operator, both sides must be true
	if node.NodeType == "AND" {
		fmt.Println("Evaluating AND")
		return EvaluateRule(node.Left, userData) && EvaluateRule(node.Right, userData)
	}

	// If it's an OR operator, either side can be true
	if node.NodeType == "OR" {
		fmt.Println("Evaluating OR")
		return EvaluateRule(node.Left, userData) || EvaluateRule(node.Right, userData)
	}

	// If it's an operand, evaluate the condition
	if node.NodeType == "operand" {
		condition := node.Value
		switch condition.Attribute {
		case "age":
			return evaluateCondition(condition, userData.Age)
		case "department":
			return evaluateCondition(condition, userData.Department)
		case "salary":
			return evaluateCondition(condition, userData.Salary)
		case "experience":
			return evaluateCondition(condition, userData.Experience)
		}
	}
	return false
}

// Helper function to evaluate a condition against a user attribute
func evaluateCondition(condition *models.Condition, userValue interface{}) bool {
	fmt.Printf("Evaluating Condition: Attribute=%s, Operator=%s, Value=%v, UserValue=%v\n", condition.Attribute, condition.Operator, condition.Value, userValue)

	switch condition.Operator {
	case ">":
		// Ensure both values are integers for comparison
		if userInt, ok := userValue.(int); ok {
			if conditionInt, ok := condition.Value.(int); ok {
				return userInt > conditionInt
			}
		}
	case "<":
		// Ensure both values are integers for comparison
		if userInt, ok := userValue.(int); ok {
			if conditionInt, ok := condition.Value.(int); ok {
				return userInt < conditionInt
			}
		}
	case "=":
		// Compare both strings or both integers
		switch conditionValue := condition.Value.(type) {
		case int:
			if userInt, ok := userValue.(int); ok {
				return userInt == conditionValue
			}
		case string:
			if userStr, ok := userValue.(string); ok {
				return userStr == conditionValue
			}
		}
	}

	return false
}
