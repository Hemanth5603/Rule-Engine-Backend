package models

type RuleNode struct {
	ID         int    `db:"id"`
	NodeType   string `db:"node_type"`
	LeftChild  *int   `db:"left_child"`
	RightChild *int   `db:"right_child"`
	Attribute  string `db:"attribute"`
	Operator   string `db:"operator"`
	Value      string `db:"value"`
}
