package parser

import "strconv"

func parseNumber(tokens []Token, i int) (node *Node, tokensConsumed int) {
	if tokens[i].Type == TokenTypeNumber {
		value, e := strconv.ParseFloat(tokens[i].Value, 64)

		if e == nil {
			var node Node = NumberNode{value}

			return &node, 1
		}
	}

	return nil, 0
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

	return nil, 0
}

func parseString(tokens []Token, i int) (node *Node, tokensConsumed int) {
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

	return nil, 0
}

// TODO this needs to be recursive
func Parse(tokens []Token) (ast *Node, ok bool) {
	check := func(node *Node, tokensConsumed int) bool {
		if node == nil {
			return false
		}

		if tokensConsumed != len(tokens) {
			return false
		}

		ast = node
		return true
	}

	if check(parseNumber(tokens, 0)) {
		return ast, true
	}

	if check(parseBoolean(tokens, 0)) {
		return ast, true
	}

	return nil, false
}
