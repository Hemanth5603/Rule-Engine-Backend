package utils

import (
	"errors"

	"github.com/hemanth5603/RuleEngineBackend/models"
)

func CombineRules(rules []*models.Node, operator string) (*models.Node, error) {
	if len(rules) == 0 {
		return nil, errors.New("no rules provided to combine")
	}

	if operator != "AND" && operator != "OR" {
		return nil, errors.New("invalid operator; must be 'AND' or 'OR'")
	}

	combined := rules[0] // Start with the first rule

	for i := 1; i < len(rules); i++ {
		combined = &models.Node{
			NodeType: operator,
			Left:     combined,
			Right:    rules[i],
		}
	}

	return combined, nil
}
