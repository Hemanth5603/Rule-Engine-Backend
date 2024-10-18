package utils

import "database/sql"

func DBfetchNodeByID(db *sql.DB, id int) (*RuleNode, error) {
	var node RuleNode
	query := `SELECT id, node_type, left_child, right_child, attribute, operator, value FROM ast_nodes WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&node.ID, &node.NodeType, &node.LeftChild, &node.RightChild, &node.Attribute, &node.Operator, &node.Value)
	if err != nil {
		return nil, err
	}
	return &node, nil
}
