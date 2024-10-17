package models

type Node struct {
	NodeType string
	Left     *Node
	Right    *Node
	Value    *Condition
}

type Condition struct {
	Attribute string
	Operator  string
	Value     interface{}
}
