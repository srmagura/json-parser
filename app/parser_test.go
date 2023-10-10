package app

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

func areAstsEqual(expected *Node, actual *Node) bool {
	// TODO
	return true
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
