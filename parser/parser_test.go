package parser

import (
	"encoding/json"
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

func areAstsEqual(expected *Node, actual *Node) bool {
	expectedJsonBytes, e1 := json.Marshal(expected)
	actualJsonBytes, e2 := json.Marshal(actual)

	if e1 != nil || e2 != nil {
		return false
	}

	expectedJson := string(expectedJsonBytes)
	actualJson := string(actualJsonBytes)

	return expectedJson == actualJson
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

	if !areAstsEqual(expected, actual) {
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

	// expected = ArrayNode{[]Node{
	// 	&NumberNode{7},
	// }}
	// assertAstsEqual(t, `[7]`, &expected)

	// expected = ArrayNode{[]Node{
	// 	BooleanNode{false},
	// 	NumberNode{2},
	// 	StringNode{"foo"},
	// }}
	// assertAstsEqual(t, `[false, 2, "foo"]`, &expected)
}
