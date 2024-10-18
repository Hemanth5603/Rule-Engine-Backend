package models

type Node struct {
	Id       int
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
