package main

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
	if tokens[i].Type != TokenTypeOpenSquareBracket {
		return
	}

	arrayNode := ArrayNode{[]*Node{}}

	for j := i + 1; j < len(tokens); {
		if tokens[j].Type == TokenTypeCloseSquareBracket {
			var nodeToReturn Node = arrayNode

			return &nodeToReturn, j - i + 1
		}

		if j != i+1 {
			if tokens[j].Type == TokenTypeComma {
				j++
			} else {
				return
			}
		}

		elementNode, elementTokensConsumed := parseAny(tokens, j)

		if elementNode == nil || elementTokensConsumed == 0 {
			return
		}

		arrayNode.Elements = append(arrayNode.Elements, elementNode)
		j += elementTokensConsumed

	}

	return
}

// Properties are only allowed to be direct children of objects
func parseProperty(tokens []Token, i int) (node *PropertyNode, tokensConsumed int) {
	j := i

	keyNode, keyTokensConsumed := parseString(tokens, j)

	if keyNode == nil || keyTokensConsumed == 0 {
		return
	}

	key := (*keyNode).(StringNode).Value
	j += keyTokensConsumed

	if j >= len(tokens) || tokens[j].Type != TokenTypeColon {
		return
	}

	j++

	valueNode, valueTokensConsumed := parseAny(tokens, j)

	j += valueTokensConsumed

	return &PropertyNode{key, valueNode}, j - i
}

func parseObject(tokens []Token, i int) (node *Node, tokensConsumed int) {
	if tokens[i].Type != TokenTypeOpenCurlyBrace {
		return
	}

	objectNode := ObjectNode{[]*PropertyNode{}}

	for j := i + 1; j < len(tokens); {
		if tokens[j].Type == TokenTypeCloseCurlyBrace {
			var nodeToReturn Node = objectNode

			return &nodeToReturn, j - i + 1
		}

		if j != i+1 {
			if tokens[j].Type == TokenTypeComma {
				j++
			} else {
				return
			}
		}

		propertyNode, propertyTokensConsumed := parseProperty(tokens, j)

		if propertyNode == nil || propertyTokensConsumed == 0 {
			return
		}

		objectNode.Properties = append(objectNode.Properties, propertyNode)
		j += propertyTokensConsumed
	}

	return
}

func parseAny(tokens []Token, i int) (node *Node, tokensConsumed int) {
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

	if check(parseObject(tokens, i)) {
		return
	}

	return
}

func Parse(tokens []Token) (ast *Node, ok bool) {
	ast, tokensConsumed := parseAny(tokens, 0)

	if ast != nil && tokensConsumed == len(tokens) {
		return ast, true
	}

	return
}
