package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/hemanth5603/RuleEngineBackend/models"
)

// CreateRule parses the rule string and creates the AST.
func CreateRule(ruleString string) (*models.Node, error) {
	tokens := strings.Fields(ruleString)

	if len(tokens) < 7 { // Minimum expected length for a valid rule
		return nil, errors.New("invalid rule format")
	}

	leftCondition := &models.Node{
		NodeType: "operand",
		Value: &models.Condition{
			Attribute: tokens[0],
			Operator:  tokens[1],
			Value:     parseValue(tokens[2]),
		},
	}

	operator := tokens[3] // AND or OR

	rightCondition := &models.Node{
		NodeType: "operand",
		Value: &models.Condition{
			Attribute: tokens[4],
			Operator:  tokens[5],
			Value:     parseValue(tokens[6]),
		},
	}

	// Create the root node with the operator
	rootNode := &models.Node{
		NodeType: operator,
		Left:     leftCondition,
		Right:    rightCondition,
	}

	return rootNode, nil
}

// parseValue converts a string value into an appropriate type (int or string).
func parseValue(value string) interface{} {
	// Trim any surrounding quotes and whitespace
	value = strings.TrimSpace(value)
	value = strings.Trim(value, "'")

	if intVal, err := strconv.Atoi(value); err == nil {
		return intVal // Return as an integer
	}
	return value // Return as a string
}
