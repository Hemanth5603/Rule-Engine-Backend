package utils

import (
	"fmt"

	"github.com/hemanth5603/RuleEngineBackend/models"
)

func EvaluateRule(node *models.Node, userData models.UserModel) bool {

	if node == nil {
		return false
	}

	if node.NodeType == "AND" {
		fmt.Println("Evaluating AND")
		return EvaluateRule(node.Left, userData) && EvaluateRule(node.Right, userData)
	}

	if node.NodeType == "OR" {
		fmt.Println("Evaluating OR")
		return EvaluateRule(node.Left, userData) || EvaluateRule(node.Right, userData)
	}
	fmt.Println(node.NodeType)

	if node.NodeType == "operand" {

		condition := node.Value
		fmt.Println(condition)
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

func evaluateCondition(condition *models.Condition, userValue interface{}) bool {
	str := fmt.Sprintf("Evaluating Condition: Attribute=%s, Operator=%s, Value=%v, UserValue=%v\n", condition.Attribute, condition.Operator, condition.Value, userValue)

	fmt.Println(str)
	switch condition.Operator {
	case ">":
		if userInt, ok := userValue.(int); ok {
			if conditionInt, ok := condition.Value.(int); ok {
				return userInt > conditionInt
			}
		}
	case "<":
		if userInt, ok := userValue.(int); ok {
			if conditionInt, ok := condition.Value.(int); ok {
				return userInt < conditionInt
			}
		}
	case "=":
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
