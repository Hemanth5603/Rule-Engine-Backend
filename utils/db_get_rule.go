package utils

import (
	"database/sql"
	"fmt"
)

type RuleNode struct {
	ID         int    `db:"id"`
	NodeType   string `db:"node_type"`
	LeftChild  *int   `db:"left_child"`
	RightChild *int   `db:"right_child"`
	Attribute  string `db:"attribute"`
	Operator   string `db:"operator"`
	Value      string `db:"value"`
}

func fetchAllNodes(db *sql.DB) ([]RuleNode, error) {
	rows, err := db.Query(`
		SELECT id, node_type, left_child, right_child, attribute, operator, value
		FROM ast_nodes
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []RuleNode
	for rows.Next() {
		var node RuleNode
		err := rows.Scan(&node.ID, &node.NodeType, &node.LeftChild, &node.RightChild, &node.Attribute, &node.Operator, &node.Value)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func createNodeMap(nodes []RuleNode) map[int]RuleNode {
	nodeMap := make(map[int]RuleNode)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}
	return nodeMap
}

func buildExpressionFromNode(node RuleNode, nodeMap map[int]RuleNode) string {
	// If it's an operand (leaf node), return the condition
	if node.NodeType == "operand" {
		return fmt.Sprintf("%s %s %s", node.Attribute, node.Operator, node.Value)
	}

	// Recursively build the left and right expressions
	var leftExpr, rightExpr string
	if node.LeftChild != nil {
		leftExpr = buildExpressionFromNode(nodeMap[*node.LeftChild], nodeMap)
	}
	if node.RightChild != nil {
		rightExpr = buildExpressionFromNode(nodeMap[*node.RightChild], nodeMap)
	}

	// Combine left and right expressions with the operator (AND/OR)
	return fmt.Sprintf("(%s %s %s)", leftExpr, node.NodeType, rightExpr)
}

func BuildExpressionsForAllNodes(db *sql.DB) ([]string, error) {
	nodes, err := fetchAllNodes(db)
	if err != nil {
		return nil, err
	}

	// Create a map for easy lookup of child nodes
	nodeMap := createNodeMap(nodes)

	// Store expressions for each root node (no parent)
	var expressions []string
	for _, node := range nodes {
		// Assuming root nodes are operators without parents
		if node.NodeType == "AND" || node.NodeType == "OR" {
			expression := buildExpressionFromNode(node, nodeMap)
			expressions = append(expressions, expression)
		}
	}
	return expressions, nil
}
