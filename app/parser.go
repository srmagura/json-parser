package app

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

	return nil, false
}
