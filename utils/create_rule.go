package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/hemanth5603/RuleEngineBackend/models"
)

func CreateRule(ruleString string) (*models.Node, error) {
	tokens := strings.Fields(ruleString)

	if len(tokens) < 7 {
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

	operator := tokens[3]

	rightCondition := &models.Node{
		NodeType: "operand",
		Value: &models.Condition{
			Attribute: tokens[4],
			Operator:  tokens[5],
			Value:     parseValue(tokens[6]),
		},
	}

	rootNode := &models.Node{
		NodeType: operator,
		Left:     leftCondition,
		Right:    rightCondition,
	}

	return rootNode, nil
}

func parseValue(value string) interface{} {

	value = strings.TrimSpace(value)
	value = strings.Trim(value, "'")

	if intVal, err := strconv.Atoi(value); err == nil {
		return intVal
	}
	return value
}
