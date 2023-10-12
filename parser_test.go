package main

import (
	"testing"
)

func assertParsingFails(t *testing.T, json string) {
	tokens, lexOk := Lex(json)

	if !lexOk {
		t.Fatal("Expected lexing to succeed, but it failed.")
	}

	_, parseOk := Parse(tokens)

	if parseOk {
		t.Fatal("Expected parsing to fail, but it succeeded.")
	}
}

func areAstsEqual(t *testing.T, expected *Node, actual *Node) bool {
	if (*expected).GetNodeType() != (*actual).GetNodeType() {
		return false
	}

	switch nodeType := (*expected).GetNodeType(); nodeType {
	case NodeTypeNumber:
		return (*expected).(NumberNode).Value == (*actual).(NumberNode).Value
	case NodeTypeBoolean:
		return (*expected).(BooleanNode).Value == (*actual).(BooleanNode).Value
	case NodeTypeString:
		return (*expected).(StringNode).Value == (*actual).(StringNode).Value
	case NodeTypeArray:
		expectedElements := (*expected).(ArrayNode).Elements
		actualElements := (*actual).(ArrayNode).Elements

		if len(expectedElements) != len(actualElements) {
			return false
		}

		for i := 0; i < len(expectedElements); i++ {
			if !areAstsEqual(t, &(expectedElements[i]), &(actualElements[i])) {
				return false
			}
		}

		return true
	}

	t.Fatal("areAstsEqual: none of the switch cases matched.")
	return false
}

func assertAstsEqual(t *testing.T, json string, expected *Node) {
	tokens, lexOk := Lex(json)

	if !lexOk {
		t.Fatal("Lexing failed.")
	}

	actual, parseOk := Parse(tokens)

	if !parseOk {
		t.Fatal("Parsing failed.")
	}

	if !areAstsEqual(t, expected, actual) {
		// If this happens, use the debugger to see how `expected` and `actual`
		// differ
		t.Fatal("ASTs are not equal.")
	}
}

func TestParseNumber(t *testing.T) {
	var expected Node = NumberNode{12.3}
	assertAstsEqual(t, "12.3", &expected)
}

func TestParseInvalidConsecutiveValues(t *testing.T) {
	assertParsingFails(t, "12.3.2")
	assertParsingFails(t, "12.3false")
	assertParsingFails(t, "12.3 false")
	assertParsingFails(t, `12."foo"`)
}

func TestParseBoolean(t *testing.T) {
	var expected Node = BooleanNode{true}
	assertAstsEqual(t, "true", &expected)

	expected = BooleanNode{false}
	assertAstsEqual(t, "false", &expected)
}

func TestParseString(t *testing.T) {
	var expected Node = StringNode{"foo bar"}
	assertAstsEqual(t, `"foo bar"`, &expected)
}

func TestParseArray(t *testing.T) {
	var expected Node = ArrayNode{[]Node{}}
	assertAstsEqual(t, `[]`, &expected)

	expected = ArrayNode{[]Node{
		NumberNode{7},
	}}
	assertAstsEqual(t, `[7]`, &expected)

	expected = ArrayNode{[]Node{
		BooleanNode{false},
		NumberNode{2},
		StringNode{"foo"},
	}}
	assertAstsEqual(t, `[false, 2, "foo"]`, &expected)
}
