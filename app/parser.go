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
	var root *Node
	i := 0

	check := func(node *Node, tokensConsumed int) bool {
		if node == nil {
			return false
		}

		// TODO
		root = node

		i += tokensConsumed
		return true
	}

	for i < len(tokens) {
		if check(parseNumber(tokens, i)) {
			continue
		}

		return nil, false
	}

	return root, true
}
