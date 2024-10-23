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

type RuleExpression struct {
	ID         int    `json:"id"`
	Expression string `json:"expression"`
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

	if node.NodeType == "operand" {
		return fmt.Sprintf("%s %s %s", node.Attribute, node.Operator, node.Value)
	}

	var leftExpr, rightExpr string
	if node.LeftChild != nil {
		leftExpr = buildExpressionFromNode(nodeMap[*node.LeftChild], nodeMap)
	}
	if node.RightChild != nil {
		rightExpr = buildExpressionFromNode(nodeMap[*node.RightChild], nodeMap)
	}

	return fmt.Sprintf("(%s %s %s)", leftExpr, node.NodeType, rightExpr)
}

func BuildExpressionsForAllNodes(db *sql.DB) ([]RuleExpression, error) {
	nodes, err := fetchAllNodes(db)
	if err != nil {
		return nil, err
	}

	// Create a map for easy lookup of child nodes
	nodeMap := createNodeMap(nodes)

	// Store expressions along with their corresponding IDs
	var expressions []RuleExpression
	for _, node := range nodes {
		// Assuming root nodes are operators without parents
		if node.NodeType == "AND" || node.NodeType == "OR" {
			expression := buildExpressionFromNode(node, nodeMap)
			expressions = append(expressions, RuleExpression{
				ID:         node.ID,
				Expression: expression,
			})
		}
	}
	return expressions, nil
}
