package utils

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/hemanth5603/RuleEngineBackend/models"
)

func DBSaveRule(db *sql.DB, node *models.Node) (int, error) {
	var nodeID int

	err := db.QueryRow(`
		INSERT INTO ast_nodes (node_type, left_child, right_child, attribute, operator, value)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`,
		node.NodeType,
		func() *int {
			if node.Left != nil {
				id, _ := DBSaveRule(db, node.Left) // Recursively save left child
				return &id
			}
			return nil
		}(),
		func() *int {
			if node.Right != nil {
				id, _ := DBSaveRule(db, node.Right) // Recursively save right child
				return &id
			}
			return nil
		}(),
		func() string {
			if node.NodeType == "operand" {
				return node.Value.Attribute //node.Value.(*models.Condition).Attribute
			}
			return ""
		}(),
		func() string {
			if node.NodeType == "operand" {
				return node.Value.Operator
			}
			return ""
		}(),
		func() string {
			if node.NodeType == "operand" {
				return parseValueToString(node.Value.Value)
			}
			return ""
		}(),
	).Scan(&nodeID)

	if err != nil {
		return 0, err
	}

	return nodeID, nil
}

// Helper function to convert value to string for insertion
func parseValueToString(value interface{}) string {
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	default:
		return fmt.Sprintf("%v", v) // Handles strings and other types
	}
}
