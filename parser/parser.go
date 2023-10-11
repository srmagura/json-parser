package parser

import (
	"strconv"
)

func parseNumber(tokens []Token, i int) (node *Node, tokensConsumed int) {
	if tokens[i].Type == TokenTypeNumber {
		value, e := strconv.ParseFloat(tokens[i].Value, 64)

		if e == nil {
			var node Node = NumberNode{value}

			return &node, 1
		}
	}

	return
}

func parseBoolean(tokens []Token, i int) (node *Node, tokensConsumed int) {
	if tokens[i].Type == TokenTypeBoolean {
		if tokens[i].Value == "true" {
			var node Node = BooleanNode{true}

			return &node, 1
		}

		if tokens[i].Value == "false" {
			var node Node = BooleanNode{false}

			return &node, 1
		}
	}

	return
}

func parseString(tokens []Token, i int) (node *Node, tokensConsumed int) {
	if tokens[i].Type == TokenTypeString {
		value := tokens[i].Value

		// Remove the leading and trailing quotes
		var node Node = StringNode{value[1 : len(value)-1]}

		return &node, 1
	}

	return
}

func parseArray(tokens []Token, i int) (node *Node, tokensConsumed int) {
	if tokens[i].Type != TokenTypeArrayStart {
		return
	}

	arrayNode := ArrayNode{[]Node{}}

	for j := i + 1; j < len(tokens); {
		if tokens[j].Type == TokenTypeArrayEnd {
			var nodeToReturn Node = arrayNode

			return &nodeToReturn, j - i + 1
		}

		elementNode, elementTokensConsumed := parseCore(tokens, j)

		if elementNode == nil || elementTokensConsumed == 0 {
			return
		}

		arrayNode.Elements = append(arrayNode.Elements, *elementNode)
		j += elementTokensConsumed
	}

	return
}

func parseCore(tokens []Token, i int) (node *Node, tokensConsumed int) {
	check := func(_node *Node, _tokensConsumed int) bool {
		if _node == nil || _tokensConsumed == 0 {
			return false
		}

		node = _node
		tokensConsumed = _tokensConsumed

		return true
	}

	if check(parseNumber(tokens, i)) {
		return
	}

	if check(parseBoolean(tokens, i)) {
		return
	}

	if check(parseString(tokens, i)) {
		return
	}

	if check(parseArray(tokens, i)) {
		return
	}

	return
}

func Parse(tokens []Token) (ast *Node, ok bool) {
	ast, tokensConsumed := parseCore(tokens, 0)

	if ast != nil && tokensConsumed == len(tokens) {
		return ast, true
	}

	return
}
