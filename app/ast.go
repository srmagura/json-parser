package app

type NodeType byte

const (
	NodeTypeNumber NodeType = iota
	NodeTypeBoolean
	NodeTypeString
	NodeTypeArray
)

type NumberNode struct {
	Value float64
}

type BooleanNode struct {
	Value bool
}

type StringNode struct {
	Value string
}

// type ArrayNode struct {
// 	Elements []Node
// }

func (n NumberNode) GetNodeType() NodeType {
	return NodeTypeNumber
}

func (n BooleanNode) GetNodeType() NodeType {
	return NodeTypeBoolean
}

func (n StringNode) GetNodeType() NodeType {
	return NodeTypeString
}

type Node interface {
	GetNodeType() NodeType
}
