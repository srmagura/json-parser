package main

type NodeType byte

const (
	NodeTypeNumber NodeType = iota
	NodeTypeBoolean
	NodeTypeString
	NodeTypeArray
	NodeTypeProperty
	NodeTypeObject
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

type ArrayNode struct {
	Elements []*Node
}

type PropertyNode struct {
	Key   string
	Value *Node
}

type ObjectNode struct {
	Properties []*PropertyNode
}

func (n NumberNode) GetNodeType() NodeType {
	return NodeTypeNumber
}

func (n BooleanNode) GetNodeType() NodeType {
	return NodeTypeBoolean
}

func (n StringNode) GetNodeType() NodeType {
	return NodeTypeString
}

func (n ArrayNode) GetNodeType() NodeType {
	return NodeTypeArray
}

func (n PropertyNode) GetNodeType() NodeType {
	return NodeTypeProperty
}

func (n ObjectNode) GetNodeType() NodeType {
	return NodeTypeObject
}

type Node interface {
	GetNodeType() NodeType
}
